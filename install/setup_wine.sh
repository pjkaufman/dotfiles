#!/usr/bin/env bash

if is_work_computer ; then
  echo "skipping wine setup"
  return
fi

if ! command -v wine &> /dev/null ; then
    wget https://dl.winehq.org/wine-builds/Release.key
    sudo apt-key add Release.key
    sudo apt-add-repository 'https://dl.winehq.org/wine-builds/ubuntu/'
fi

install_apt_package "wine" "--install-recommends winehq-stable"

# fix Kindle for PC 1.17 network connectivity as per https://askubuntu.com/a/1352999
# regenerate certs and add line to trusted sources as per https://askubuntu.com/a/1352999
cert_text="mozilla/VeriSign_Class_3_Public_Primary_Certification_Authority_-_G5.crt" 
cert_conf="/etc/ca-certificates.conf"

if ! grep -q "$cert_text" "$cert_conf" ; then
  echo "$cert_text" >> "$cert_conf"
fi

cert_file=/usr/share/ca-certificates/mozilla/VeriSign_Class_3_Public_Primary_Certification_Authority_-_G5.crt
if [ ! -f "$cert_file" ]; then 
  sudo cp "$HOME/dotfiles/wine/certificate" "$cert_file"
  sudo update-ca-certificates
fi

expected_cert_file="/etc/ssl/certs/b204d74a.0"
if [ ! -f "$expected_cert_file" ]; then
  echo "Something is not right about the cert generation as file '$expected_cert_file' was not generated."
  exit 1;
fi
