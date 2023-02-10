#!/bin/bash

# helper functions

# install_apt_package installs an apt package if it does not already exist
install_apt_package() {

  if ! command -v $1 &> /dev/null
  then
    echo "installing $1"
    
    if [ -z $2 ]
    then
      sudo apt install -y $2
    else 
      sudo apt install -y $1
    fi
  else
    echo "$1 is already installed"
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

  if ! command -v $1 &> /dev/null
  then
    echo "installing $1"
    pip3 install --user $1  
  else
    echo "$1 is already installed"
  fi
}

# go_install_package installs a go package if it is not currently installed
# $1 is the full installation value for go install to use 
# $2 is the short name of the package to use in output for the script and the actual command name
go_install_package() {

  if ! command -v $2 &> /dev/null
  then
    echo "installing $2"
    go install $1@latest 
  else
    echo "$2 is already installed"
  fi
}

# npm_install_package installs an npm package if it is not currently installed
# $1 is the npm package to install globally
npm_install_package() {

  if ! command -v $1 &> /dev/null
  then
    echo "installing $1"
    npm install -g $1 
  else
    echo "$1 is already installed"
  fi
}

# cargo_install_package installs a cargo package if it is not currently installed
# $1 is the name of the cargo package to install
cargo_install_package() {

  if ! command -v $1 &> /dev/null
  then
    echo "installing $1"
    cargo install $1 
  else
    echo "$1 is already installed"
  fi
}

# handle_flatpak_installations determines whether or not to install the flatpaks
# and installs any that are missing
# $1 is whether or not the computer is a work computer
handle_flatpak_installations() {

  if $1
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
    echo 'export COMPUTER_TYPE=personal' >> ~/.bashrc 
  else
    is_work_computer=true
    echo 'export COMPUTER_TYPE=work' >> ~/.bashrc 
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

install_apt_package "git"
install_apt_package "tmux"
install_apt_package "grep"
install_apt_package "curl"
install_apt_package "btop" # task manager equivalent
install_apt_package "python3"
install_apt_package "pip3" "python3-pip"
install_apt_package "cargo"

if ! $is_work_computer
then
  install_apt_package "imgp" # image compression
  install_apt_package "pandoc" # document conversion
  install_apt_package "flameshot" # screenshots
  install_apt_package "kitty" # terminal
fi

# cargo packages

setup_header_text "cargo packages:"

cargo_install_package "topgrade"

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

# TODO: special package managers like nvm and gvm

setup_header_text "gvm and nvm install:"

if ! command -v gvm &> /dev/null
then
  echo "installing gvm"
  bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
else
  echo "gvm is already installed"
fi

if [ -z ${NVM_DIR} ]
then
  echo "installing nvm"
  curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
else
  echo "nvm is already installed"
fi

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

if $is_work_computer
then
  go_install_package "github.com/yoheimuta/protolint/cmd/protolint" "protolint"
  go_install_package "golang.org/x/tools/cmd/goimports " "golangci-lint"
fi

go_install_package "golang.org/x/tools/cmd/gofmt" "gofmt"
go_install_package "golang.org/x/tools/cmd/goimports" "goimports"

npm_install_package "eslint"

cargo_install_package "stylua"

# TODO: handle google_java_format install

# TODO: move scripts to bin

# TODO: setup config symlinks

echo ""
echo "environment setup complete"

