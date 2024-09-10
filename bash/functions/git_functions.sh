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
  git checkout $(git symbolic-ref --short refs/remotes/origin/HEAD | cut -d "/" -f2)
}

# gs searchs for a branch to switch to in the current repo using fzf  
function gs() {
  local branch
  branch=$(git branch -a | grep -v '/HEAD\s' | grep -v '^\*' | sed 's/remotes\/origin\///' | sed 's/^[ \t]*//' | sort -u | fzf --exit-0)
  
  if [[ -n $branch ]]; then
    if git rev-parse --verify "$branch" >/dev/null 2>&1; then
      git checkout "$branch"
    else
      git checkout --track "origin/$branch"
    fi
  else
    echo "No branch selected. Staying on the current branch."
  fi
}
