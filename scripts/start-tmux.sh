#!/bin/bash

session="prod"

tmux has-session -t $session 2>/dev/null
if [ "$?" -eq 1 ] ; then
    tmux new-session -d -s $session
    window=0
    # set title
    tmux send-keys -t $session:$window "printf '\033]2;%s\033\\\\' 'btop'" C-m
    tmux send-keys -t $session:$window 'btop' C-m
    # split vertically
    tmux split-window -h -t 0
    tmux send-keys -t $session:$window "printf '\033]2;%s\033\\\\' 'productivity'" C-m
    tmux send-keys -t $session:$window "c" C-m
fi
tmux attach-session -t $session
