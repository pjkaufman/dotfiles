#!/usr/bin/env sh

# wsl aliases

function is_wsl() {
  local is_wsl=$(uname -r | grep -i "Microsoft")
  if [[ -z "$is_wsl" ]]; then
    return 0
  fi

  return 1
}

if is_wsl ; then
  return
fi

# This is specific to WSL 2. If the WSL 2 VM goes rogue and decides not to free
# up memory, this command will free your memory after about 20-30 seconds.
#   Details: https://github.com/microsoft/WSL/issues/4166#issuecomment-628493643
# This is an alias because creating the bash function for it fails to run properly
alias drop_cache="sudo sh -c \"echo 3 >'/proc/sys/vm/drop_caches' && swapoff -a && swapon -a && printf '\n%s\n' 'Ram-cache and Swap Cleared'\""
