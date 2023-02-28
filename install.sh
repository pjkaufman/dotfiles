#!/bin/bash

# helper functions

# check_if_command_exists_and_run_install_command_otherwise if a command does not exist, it runs the install command provided
# $1 is the command to verify exists 
# $2 is the command to install the 
check_if_command_exists_and_run_install_command_otherwise() {
  if ! command -v $1 &> /dev/null
  then
    echo "installing $1"
    eval " $2" 
  else
    echo "$1 is already installed"
  fi
}

# install_apt_package installs an apt package if it does not already exist
# $1 is the name of the command to verify exists 
# $2 is the actual name of the package to install for the command if the name is different from the command
install_apt_package() {
  if [ -z $2 ]
  then
    check_if_command_exists_and_run_install_command_otherwise $1 "sudo apt install -y $1"
  else
    check_if_command_exists_and_run_install_command_otherwise $1 "sudo apt install -y $2"
  fi
}

# install_flatpak_package installs a flatpak package if the flatpak name is not present
# in the flatpak list of installed packages
# $1 is the name of the flatpak to install
# $2 is the actual package name to install (i.e. the one with all of the periods in it)
install_flatpak_package() {
  grep_output=`flatpak list | grep $2`
  if [ -z grep_output ]
  then
      echo "installing $1"
      flatpak install --user $2
  else
    echo "$1 is already installed"
  fi
}

# add_ppa_and_install_package adds the specified PPA if it does not exist already
# $1 is the name of the package with a slash and the type of stability of the PPA (i.e. syncthing/stable)
# $2 is the name of the package to install once the PPA has been added
add_ppa_and_install_package() {
  grep_output=`apt-cache policy| grep $1`
  if [ -z grep_output ]
  then
    echo "adding $1 PPA and installing $2"

    sudo add-apt-repository ppa:$1
    sudo apt update && sudo apt install -y $2
  else
    echo "$1 PPA already added"
  fi
}

# pip_install_package installs pip packages if it is not currently installed
# $1 is the name of the pip package to install if it is not currently present
pip_install_package() {
  check_if_command_exists_and_run_install_command_otherwise $1 "pip3 install --user $1"  
}

# go_install_package installs a go package if it is not currently installed
# $1 is the full installation value for go install to use 
# $2 is the short name of the package to use in output for the script and the actual command name
go_install_package() {
  check_if_command_exists_and_run_install_command_otherwise $2 "go install $1@latest" 
}

# npm_install_package installs an npm package if it is not currently installed
# $1 is the npm package to install globally
npm_install_package() {
  check_if_command_exists_and_run_install_command_otherwise $1 "npm install -g $1" 
}

# cargo_install_package installs a cargo package if it is not currently installed
# $1 is the name of the cargo package to install
cargo_install_package() {
  check_if_command_exists_and_run_install_command_otherwise $1 "cargo install $1" 
}

# handle_flatpak_installations determines whether or not to install the flatpaks
# and installs any that are missing
# $1 is whether or not the computer is a work computer
handle_flatpak_installations() {
  if ! $1
  then
    if ! command -v flatpak &> /dev/null
    then
      echo "Flatpak not installed. Please install it."
    else
      install_flatpak_package "Brave Browser" "com.brave.Browser"
      install_flatpak_package "Minecraft" "com.mojang.Minecraft"
      install_flatpak_package "GnuCash" "com.gnucash.GnuCash"
      install_flatpak_package "Sigil" "com.sigil_ebook.Sigil"
      install_flatpak_package "Calibre" "com.calibre_ebook.calibre"
      install_flatpak_package "Obsidian" "md.obsidian.Obsidian"
      install_flatpak_package "Only Office" "org.onlyoffice.desktopeditors"
    fi
  else
    echo "Skipping flatpak installation"
  fi
}

# prepend_text_to_file takes the specified text and adds it to the start of the file specified
# $1 is the text to add to the start of the file
# $2 is the path for the file to put the text at the start of
prepend_text_to_file() {
  echo "$1" > /tmp/tmpfile.$$
  cat "$2" >> /tmp/tmpfile.$$
  mv /tmp/tmpfile.$$ "$2"
}

ensure_file_symlink_is_in_place() {
  if [ -L $2 ] ; then
    if [ -e $2 ] ; then
      echo "'$2' is already symlinked"
    else
      echo "'$2' is a broken symlink"
    fi
  elif [ -e $2 ] ; then
    echo "'$2' exists, but is not symlinked"
    mv "$2" "$2.bak"
  else
    echo "'$2' does not exist"
  fi
  
  ln -sf "$1" "$2"  
}


ensure_folder_symlink_is_in_place() {
  if [ -L $2 ] ; then
    if [ -d $2 ] ; then
      echo "'$2' is already symlinked, please check that it is the correct symlink"
      return
    else
      echo "'$2' is a broken symlink"
      ln -s "$1" "$2"
    fi
  elif [ -d $2 ] ; then
    echo "'$2' exists, but is not symlinked (implementation needed)"
    return
  else
    echo "'$2' does not exist"
      ln -s "$1" "$2"
  fi
}

setup_header_text() {
  echo ""
  echo "==============================="
  echo "$1"
  echo "==============================="
  echo ""
}

# actual setup

echo "starting environment setup"

# determine whether or not this env is work or personal

setup_header_text "get computer type:"

is_work_computer=false
if [ -z "${COMPUTER_TYPE}" ]
then
  read -p 'Is this a personal computer? [y]es or [n]o: ' response_char

  if [ response_char = "y" ]
  then
    is_work_computer=false
    prepend_text_to_file 'export COMPUTER_TYPE=personal' ~/.local_extra
  else
    is_work_computer=true
    prepend_text_to_file 'export COMPUTER_TYPE=work' ~/.local_extra
  fi

  echo "Please make sure to run source your profile after the install."
else
  if [ ${COMPUTER_TYPE} = "personal" ]
  then
    is_work_computer=false
  else
    is_work_computer=true
  fi

  echo "The computer is a ${COMPUTER_TYPE} one."
fi

# apt packages

setup_header_text "apt packages:"

declare -A apt_packages_to_install=( ["git"]="git" ["tmux"]="tmux" ["grep"]="grep" ["curl"]="curl" ["btop"]="btop" ["python3"]="python3" ["pip3"]="python3-pip" ["rename"]="rename" ["cargo"]="cargo")
for pkg in "${!apt_packages_to_install[@]}"; do install_apt_package "$pkg" "${apt_packages_to_install[$pkg]}"; done

if ! $is_work_computer
then
  personal_apt_packages_to_install=(
    "imgp" # image compression
    "pandoc" # document conversion
    "flameshot" # screenshots
    "kitty" # terminal
    "evince" # pdf editor and viewer
  )
  for pkg in "${personal_apt_packages_to_install[@]}"; do install_apt_package "$pkg"; done
fi

# cargo packages

setup_header_text "cargo packages:"

cargo_packages_to_install=("topgrade" "cargo-update")
for pkg in "${cargo_packages_to_install[@]}"; do cargo_install_package "$pkg"; done

# PPA additions

setup_header_text "PPA additions"

if ! $is_work_computer
then 
  add_ppa_and_install_package "syncthing/stable" "syncthing"
fi

add_ppa_and_install_package "neovim-ppa/unstable" "neovim"

# flatpak packages

setup_header_text "Flatpak packages:"

handle_flatpak_installations $is_work_computer

# special package managers like nvm and gvm

setup_header_text "gvm and nvm install:"

check_if_command_exists_and_run_install_command_otherwise "gvm" "curl -sSL https://github.com/soulteary/gvm/raw/master/binscripts/gvm-installer | bash"

# nvm is special and loads its command via autocompletion and checking the created variable
# is more reliable than checking if the method exists
if [ -z ${NVM_DIR} ]
then
  echo "installing nvm"
  curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
else
  echo "nvm is already installed"
fi

# TODO: add logic for wkhtml to pdf

# TODO: add logic around ssh agent for github

# i3 setup

setup_header_text "i3 setup:"

if ! $is_work_computer
then 
  install_apt_package "i3"
  install_apt_package "scrot" # screenshots
  install_apt_package "light" # brightness changes
  install_apt_package "feh" # background image
  install_apt_package "i3lock" # lockscreen setup
  install_apt_package "i3status" # status info
  install_apt_package "pactl" # sound changes
  install_apt_package "dmenu" # app selector
  pip_install_package "bumblebee-status" # status bar
else
  echo "skipping i3 setup"
fi

# neovim setup

setup_header_text "Neovim setup:"

# null-ls lsp formatters and diagnostics 
pip_install_package "codespell"
pip_install_package "black"
pip_install_package "flake8"
pip_install_package "beautysh"

if $is_work_computer
then
  go_install_package "github.com/yoheimuta/protolint/cmd/protolint" "protolint"
  go_install_package "golang.org/x/tools/cmd/goimports" "golangci-lint"
fi

# go_install_package "golang.org/x/tools/cmd/gofmt" "gofmt" # is a part of go
go_install_package "golang.org/x/tools/cmd/goimports" "goimports"

npm_install_package "eslint"

cargo_install_package "stylua"

# TODO: handle google_java_format install

# setup config symlinks

setup_header_text "Symlink setup:"

ensure_folder_symlink_is_in_place "$HOME/dotfiles/nvim" "$HOME/.config/nvim"

declare -A file_symlink_info=( 
  ["$HOME/dotfiles/git/.gitconfig"]="$HOME/.gitconfig" 
  ["$HOME/dotfiles/.shellrc/bash_profile"]="$HOME/.bash_profile" 
  ["$HOME/dotfiles/.shellrc/bashrc"]="$HOME/.bashrc" 
  ["$HOME/dotfiles/.shellrc/hushlogin"]="$HOME/.hushlogin"
  ["$HOME/dotfiles/topgrade/topgrade.toml"]="$HOME/.config/topgrade.toml"
)

for file in "${!file_symlink_info[@]}"; do ensure_file_symlink_is_in_place "$file" "${file_symlink_info[$file]}"; done

if ! $is_work_computer
then
  ensure_file_symlink_is_in_place "$HOME/dotfiles/kitty/kitty.conf" "$HOME/.config/kitty/kitty.conf"
  ensure_file_symlink_is_in_place "$HOME/dotfiles/i3/config" "$HOME/.config/i3/config"
fi

echo ""
echo "environment setup complete"
