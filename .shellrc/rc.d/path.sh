# Prepend new items to path (if directory exists)

prepend_path "/bin"
prepend_path "/usr/bin"
prepend_path "/usr/local/bin"
prepend_path "$HOME/bin"
prepend_path "$HOME/.local/bin"
prepend_path "$CARGO_HOME/bin"
prepend_path "/sbin"
prepend_path "/usr/sbin"
prepend_path "$GOROOT/bin"

# Remove duplicates (preserving prepended items)
# Source: http://unix.stackexchange.com/a/40755

PATH=$(echo -n $PATH | awk -v RS=: '{ if (!arr[$0]++) {printf("%s%s",!ln++?"":":",$0)}}')

# Wrap up

export PATH
