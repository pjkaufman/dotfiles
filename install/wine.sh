#!/bin/bash

if is_work_computer ; then
  echo "skipping wine setup"
  return
fi

wget -qO- https://dl.winehq.org/wine-builds/winehq.key | sudo apt-key add -
sudo apt-add-repository "deb http://dl.winehq.org/wine-builds/ubuntu/ $(lsb_release -cs) main"
sudo apt install --install-recommends winehq-stable

# fix Kindle for PC 1.17 network connectivity as per https://askubuntu.com/a/1352999

// TODO: add mozilla/VeriSign_Class_3_Public_Primary_Certification_Authority_-_G5.crt in the file /etc/ca-certificates.conf if missing

cert_file=/usr/share/ca-certificates/mozilla/VeriSign_Class_3_Public_Primary_Certification_Authority_-_G5.crt
if [ ! -f $cert_file ]; then 
  sudo cp $HOME/dotfiles/wine/certificate $cert_file
  sudo update-ca-certificates
fi

# TODO: regenerate certs and add line to trusted sources as per https://askubuntu.com/a/1352999
