# Moving around
alias ..='cd ..'
alias ...='cd ../../'
alias ....='cd ../../../'
alias .....='cd ../../../../'

# git
alias gu='git undo'

# ls 
alias ll='ls -alF'
alias la='ls -A'
alias l='ls -CF'

# make it so that wget can work without needing to write its config to the home dir
alias wget=wget --hsts-file="$XDG_DATA_HOME/wget-hsts"
