function c() {
  clear && clear
}

function gh() {
  history | grep "$@"
}

function python() {
  command python3 "$@"
}

function update() {
  topgrade
}

# tmux aliases
function starttmux() {
  . "$HOME/dotfiles/scripts/start-tmux.sh"
}

function killsess() {
  tmux kill-session -t "$@"
}

function tls() {
  tmux ls
}

# personal computer aliases
if [ ${COMPUTER_TYPE} = "personal" ]
then
  function hibernate() {
    sudo systemctl hibernate && echo "hello"
  }

  # allows for easy running of obsidian via terminal
  function obsidian() {
    flatpak run md.obsidian.Obsidian &
  }
  # allows for easy running of brave via terminal
  function brave() {
    flatpak run com.brave.Browser &
  }

  # allows for easy running of GnuCash via terminal
  function gnucash() {
    flatpak run org.gnucash.GnuCash &
  }

  # allows for easy running of Minecraft via terminal
  function minecraft() {
    flatpak run com.mojang.Minecraft &
  }

  # allows for easy running of Only Office via terminal
  function office() {
    flatpak run org.onlyoffice.desktopeditors &
  }

  # allows for easy running of Sigil via terminal
  function sigil() {
    flatpak run com.sigil_ebook.Sigil &
  }

  # allows for easy running of Calibre via terminal
  function calibre() {
    flatpak run com.calibre_ebook.calibre &
  }

  # enable the use of brightness since it is locked by admin permissions by default and I need to modify it using user permission
  function enablebright() {
    sudo chmod a+wr /sys/class/backlight/amdgpu_bl0/brightness
  }

  function compressepub() {
    . "$HOME/dotfiles/scripts/compress-epub.sh"
  }
fi
