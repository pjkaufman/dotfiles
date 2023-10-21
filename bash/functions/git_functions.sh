#!/usr/bin/env bash

function undo() {
  git-helper undo
}

function resetDE() {
  git-helper submodule reset
}
