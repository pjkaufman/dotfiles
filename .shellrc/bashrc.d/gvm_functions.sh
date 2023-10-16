#!/usr/bin/env bash

# gvm aliases

# gvmforceuse makes sure that the go version specified is installed before trying to use the version and
# $1 is the go version number to use (i.e. 1.19, 1.20, etc.)
gvmforceuse() {
  gvm_version_grep_output=`gvm list | grep "go$1"`
  if [[ -z "$gvm_version_grep_output" ]]
  then
    echo "installing go$1"
    gvm install "go$1" -B
  fi

  gvm use "go$1"
}
