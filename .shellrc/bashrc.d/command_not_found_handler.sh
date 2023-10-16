#!/usr/bin/env bash

#!/usr/bin/env bash

command_not_found_handle()
{
  if [[ $1 =~ .*.pdf || $1 =~ .*.PDF ]]; then
    evince "$1"
  elif [[ $1 =~ .*.jar || $1 =~ .*.JAR ]]; then
   java -jar "$1"
  elif [[ $1 =~ .*.html || $1 =~ .*.HTML || $1 =~ .*.HTM || $1 =~ .*.htm ]]; then
   brave "$1"
  elif [[ $1 =~ .*.zip || $1 =~ .*.ZIP || $1 =~ .*.war || $1 =~ .*.WAR ]]; then
    unzip -l "$1"
  elif [[ $1 =~ .*.gz || $1 =~ .*.tgz || $1 =~ .*.TGZ ]]; then
    tar -tf "$1"
  else
    # based on https://askubuntu.com/a/1464110
    printf "%s: command not found\n" "$1" >&2
    return 127
  fi
}
