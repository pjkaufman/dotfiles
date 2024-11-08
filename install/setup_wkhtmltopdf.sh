#!/usr/bin/env bash

if is_work_computer; then
  echo "skipping wkhtmltopdf install"
  return
fi

install_apt_package "pandoc"     # document conversion
install_apt_package "weasyprint" # html to pdf
