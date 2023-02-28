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
