alias c="clear && clear"
alias gh="history | grep"
alias python="python3"
alias update="topgrade"

# personal computer aliases
if [ ${COMPUTER_TYPE} = "personal" ]
then
  alias hibernate="sudo systemctl hibernate"
  # allows for easy running of obsidian via terminal
  alias obsidian="flatpak run md.obsidian.Obsidian &"
  # allows for easy running of brave via terminal
  alias brave="flatpak run com.brave.Browser &"
  # allows for easy running of GnuCash via terminal
  alias gnucash="flatpak run org.gnucash.GnuCash &"
  # allows for easy running of Minecraft via terminal
  alias minecraft="flatpak run com.mojang.Minecraft &"
  # allows for easy running of Only Office via terminal
  alias office="flatpak run org.onlyoffice.desktopeditors &"
  # allows for easy running of Sigil via terminal
  alias sigil="flatpak run com.sigil_ebook.Sigil &"
  alias calibre="flatpak run com.calibre_ebook.calibre &"
 # enable the use of brightness since it is locked by admin permissions by default and I need to modify it using user permission
  alias enablebright="sudo chmod a+wr /sys/class/backlight/amdgpu_bl0/brightness"
fi
