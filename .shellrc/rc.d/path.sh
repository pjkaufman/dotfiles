# Prepend new items to path (if directory exists)

prepend-path "/bin"
prepend-path "/usr/bin"
prepend-path "/usr/local/bin"
prepend-path "$HOME/bin"
prepend-path "$HOME/.local/bin"
# prepend-path "/usr/local/go/bin"
prepend-path "$HOME/.cargo/bin"
prepend-path "/sbin"
prepend-path "/usr/sbin"
prepend-path "$GOROOT/bin"

# Remove duplicates (preserving prepended items)
# Source: http://unix.stackexchange.com/a/40755

PATH=$(echo -n $PATH | awk -v RS=: '{ if (!arr[$0]++) {printf("%s%s",!ln++?"":":",$0)}}')

# Wrap up

export PATH
