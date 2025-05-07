#!/usr/bin/env bash

# gb is meant to git branch (i.e. create a new git branch and check it out)
function gb() {
  git switch -c "$@"
}

# gl goes to the last branch that was checkd out
function gl() {
  git checkout -
}

# gm goes to the master branch of the repo
function gm() {
  git checkout "$(git symbolic-ref --short refs/remotes/origin/HEAD | cut -d "/" -f2)"
}

# gs searchs for a branch to switch to in the current repo using fzf
function gs() {
  local branch
  branch=$(git branch -a | grep -v '/HEAD\s' | grep -v '^\*' | sed 's/remotes\/origin\///' | sed 's/^[ \t]*//' | sort -u | fzf --exit-0)

  if [[ -n $branch ]]; then
    if git rev-parse --verify "$branch" > /dev/null 2>&1; then
      git checkout "$branch"
    else
      git checkout --track "origin/$branch"
    fi
  else
    echo "No branch selected. Staying on the current branch."
  fi
}

# handle an issue with git_ps1 not 100% working with submodules on
# the windows side of a WSL2 setup
function __git_ps1_windows_mount() {
  local format_str="$1"
  # Default format if none provided
  if [ -z "$format_str" ]; then
    format_str=" (%s)"
  fi

  local path
  path=$(pwd)
  if iswsl && [[ "$path" =~ /mnt/[a-z]/ ]] && git rev-parse --is-inside-work-tree > /dev/null 2>&1; then
    # We're in a Windows mount path
    # Use git.exe directly to get branch info
    local branch
    branch=$(git.exe branch --show-current 2> /dev/null)
    if [ -n "$branch" ]; then
      # Replace %s with the actual branch name
      echo "${format_str//%s/$branch}"
    fi
  else
    # Use normal __git_ps1 for WSL paths
    __git_ps1 "$format_str"
  fi
}
