#!/bin/bash

# based on https://gist.github.com/kai101/99d57462f2459245d28b4f5ea51aa7d0

if is_work_computer; then
  echo "skipping wkhtmltopdf install"
  return
fi

if ! command -v wkhtmltopdf &> /dev/null
then
  wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz
  tar xvf wkhtmltox-0.12.4_linux-generic-amd64.tar.xz 
  sudo mv wkhtmltox/bin/wkhtmlto* $HOME/.local/bin 
fi

install_apt_package "openssl" 
install_apt_package "libssl-dev" 
install_apt_package "libxrender-dev"
install_apt_package "libx11-dev" 
install_apt_package "libxext-dev" 
install_apt_package "libfontconfig1-dev"
install_apt_package "libfreetype6-dev"
install_apt_package "fontconfig"
install_apt_package "pandoc" # document conversion
