#!/bin/bash

# Add to path
prepend_path() {
  [ -d $1 ] && PATH="$1:$PATH"
}

reload() {
  [ -f "$HOME/.bash_profile" ] && source "$HOME/.bash_profile"
}

c() {
  clear && clear
}

gh() {
  history | grep "$@"
}

python() {
  command python3 "$@"
}

update() {
  command topgrade
}

rn() {
  if [[ -z $4 ]]; then
    rename "s/$1/$2/" $3
    return;
  fi

  rename -n "s/$1/$2/" $3
}

# keyboard setup

# set keyboard layout to English international to allow alt+character to have accented or Spanish characters
eson(){
  setxkbmap -layout us -variant intl
}

# set keyboard layout to regular English to make it easier to program and do other things from the cli
esoff(){
  setxkbmap -layout us
}

# scan the computer for viruses and other issues
scan() {
  sudo rkhunter --check --sk --rwo
}