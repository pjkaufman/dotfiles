#!/bin/bash

# helper functions

# install_apt_package installs an apt package if it does not already exist
install_apt_package() {

  if ! command -v $1 &> /dev/null
  then
    echo "installing $1"
    
    if [ -z $2 ]
    then
      sudo apt install $2
    else
      sudo apt install $1
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

# add_ppa_and_install_package adds the specified PPA if 
add_ppa_and_install_package() {

}

setup_header_text() {
  echo "$1"
  echo ""
}

# actual setup
echo "starting environment setup"
echo ""

# TODO: determine whether or not this env is work or personal

setup_header_text "installing apt packages:"

# apt packages

install_apt_package "git"
install_apt_package "tmux"
install_apt_package "grep"
install_apt_package "btop"
install_apt_package "python3"
install_apt_package "pip3" "python3-pip"

# TODO: PPA additions

# flatpak packages

setup_header_text "installing flatpak packages:"

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


# TODO: special package managers like nvm and gvm

# TODO: i3 install

# TODO: move scripts to bin

# TODO: setup config symlinks

echo ""
echo "environment setup complete"

