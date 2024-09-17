#!/usr/bin/env bash

# package manager installation methods 

# check_if_command_exists_and_run_install_command_otherwise if a command does not exist, it runs the install command provided
# $1 is the command to verify exists 
# $2 is the command to install the 
check_if_command_exists_and_run_install_command_otherwise() {
  if ! command -v "$1" &> /dev/null
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
  if [ -z "$2" ]
  then
    check_if_command_exists_and_run_install_command_otherwise "$1" "sudo apt install -y $1"
  else
    check_if_command_exists_and_run_install_command_otherwise "$1" "sudo apt install -y $2"
  fi
}

# install_flatpak_package installs a flatpak package if the flatpak name is not present
# in the flatpak list of installed packages
# $1 is the name of the flatpak to install
# $2 is the actual package name to install (i.e. the one with all of the periods in it)
install_flatpak_package() {
  grep_output=$(flatpak list | grep "$2")
  if [ -z "$grep_output" ]
  then
      echo "installing $1"
      flatpak install --user "$2"
  else
    echo "$1 is already installed"
  fi
}

# install_apt_package_by_package_name_only makes sure that the provided package name
# is installed and installs it if it is not
# $1 is the name of the package to make sure to install
install_apt_package_by_package_name_only() {
  grep_output=$(dpkg -s "$1" | grep "installed")
  if [ -z "$grep_output" ]
  then
      echo "installing $1"
      sudo apt install -y "$1"
  else
    echo "$1 is already installed"
  fi
}

# add_ppa_and_install_package adds the specified PPA if it does not exist already
# $1 is the name of the package with a slash and the type of stability of the PPA (i.e. syncthing/stable)
# $2 is the name of the package to install once the PPA has been added
add_ppa_and_install_package() {
  if ! apt-cache policy | grep -q "$1" ; then
    echo "adding $1 PPA and installing $2"

    sudo add-apt-repository "ppa:$1"
    sudo apt update && sudo apt install -y "$2"
  else
    echo "$1 PPA already added"
  fi
}

# pip_install_package installs pip packages if it is not currently installed
# $1 is the name of the pip package to install if it is not currently present
pip_install_package() {
  check_if_command_exists_and_run_install_command_otherwise "$1" "pipx install $1"  
}

# go_install_package installs a go package if it is not currently installed
# $1 is the full installation value for go install to use 
# $2 is the short name of the package to use in output for the script and the actual command name
go_install_package() {
  check_if_command_exists_and_run_install_command_otherwise "$2" "go install $1@latest" 
}

# npm_install_package installs an npm package if it is not currently installed
# $1 is the npm package to install globally
npm_install_package() {
  check_if_command_exists_and_run_install_command_otherwise "$1" "npm install -g $1" 
}
