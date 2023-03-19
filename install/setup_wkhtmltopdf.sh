#!/bin/bash

# based on https://gist.github.com/kai101/99d57462f2459245d28b4f5ea51aa7d0

if is_work_computer; then
  echo "skipping wkhtmltopdf install"
  return
fi

wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz
tar xvf wkhtmltox-0.12.4_linux-generic-amd64.tar.xz 
sudo mv wkhtmltox/bin/wkhtmlto* /usr/local/bin 
sudo apt-get install -y openssl libssl-dev libxrender-dev libx11-dev libxext-dev libfontconfig1-dev libfreetype6-dev fontconfig
