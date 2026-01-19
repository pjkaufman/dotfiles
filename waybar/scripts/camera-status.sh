#!/usr/bin/env bash

# check if any cameras are present. If they are, then we are going to display a checkmark otherwise we will display an x.
if ! compgen -G "/dev/video*" > /dev/null; then
  echo '{"text":"✗","tooltip":"No camera found","class":"no-camera"}'
else
  echo '{"text":"✓","tooltip":"Camera found","class":"camera-found"}'
fi
