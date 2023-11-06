#!/usr/bin/env bash

# gets cheatsheet info for the language or command specified from the languages and commands list
# based on https://github.com/ThePrimeagen/.dotfiles/blob/602019e902634188ab06ea31251c01c1a43d1621/bin/.local/scripts/tmux-cht.sh

if ! command -v fzf &> /dev/null; then
  echo "fzf needs to be installed to run this script"
  exit 1;
fi

if ! command -v tmux &> /dev/null; then
  echo "tmux needs to be installed to run this script"
  exit 1;
fi

selected=`cat "$DOTFILES/tmux/tmux-cht-languages" "$DOTFILES/tmux/tmux-cht-command" | fzf`
if [[ -z $selected ]]; then
    exit 0
fi

read -p "Enter Query: " query

if grep -qs "$selected" "$DOTFILES/tmux/tmux-cht-languages"; then
    query=`echo $query | tr ' ' '+'`
    tmux neww bash -c "echo \"curl cht.sh/$selected/$query/\" & curl cht.sh/$selected/$query & while [ : ]; do sleep 1; done"
else
    tmux neww bash -c "curl -s cht.sh/$selected~$query | less"
fi