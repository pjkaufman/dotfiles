#!/usr/bin/env bash

undo() {
  git-helper undo
}

resetDE() {
  git-helper submodule reset
}
