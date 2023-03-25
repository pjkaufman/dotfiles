#!/bin/bash

# nvm aliases

# nvmforceuse makes sure that the node version specified is installed before trying to use the version installing it if it is not present
# $1 is the node version number to use (i.e. 16.0, 18.0, etc.)
nvmforceuse() {
  nvm_version_grep_output=`nvm list | grep "v$1"`
  if [[ -z "$nvm_version_grep_output" ]]
  then
    echo "installing v$1"
    nvm install "$1"
  fi

  nvm use "$1"
}
