#!/bin/bash

# Add to path
prepend-path() {
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
  rename 's/$1/$2/s' "$3"
}

# scan the computer for viruses and other issues
scan() {
  sudo rkhunter --check --sk --rwo
}
