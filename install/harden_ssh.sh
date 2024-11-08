#!/usr/bin/env bash

# make sure bash fails if any error happens
set -e

# originally based on https://gist.github.com/dimaskiddo/18c1c4ca71a73828c57189aba5ec5d8d

sshd_config=/etc/ssh/sshd_config

# replaces the provided sshd config values with the desired default or adds it if it is missing
# $1: name of config value
# $2: name and value
function replaceSshdConfigValue() {
  # shellcheck disable=SC2046
  if [ $(grep -c "$1" "$sshd_config") -eq 0 ]; then
    echo "$2" | sudo tee -a "$sshd_config"
  else
    sudo sed -i -e "1,/#$1 [a-zA-Z0-9]*/s/#$1 [a-zA-Z0-9]*/$2/" "$sshd_config"
    sudo sed -i -e "1,/$1 [a-zA-Z0-9]*/s/$1 [a-zA-Z0-9]*/$2/" "$sshd_config"
  fi
}

# Back-up current configuration file
sudo cp "$sshd_config" "$sshd_config.backup"

# Change value PermitUserEnvironment to no
replaceSshdConfigValue PermitUserEnvironment "PermitUserEnvironment no"

# Change value PermitEmptyPasswords to No.
replaceSshdConfigValue PermitEmptyPasswords "PermitEmptyPasswords no"

# Change value MaxAuthTries to 3
replaceSshdConfigValue MaxAuthTries "MaxAuthTries 3"

# Change value LoginGraceTime to 20.
replaceSshdConfigValue LoginGraceTime "LoginGraceTime 20"

# Disable several auth types that are not used
replaceSshdConfigValue ChallengeResponseAuthentication "ChallengeResponseAuthentication no"
replaceSshdConfigValue KerberosAuthentication "KerberosAuthentication no"
replaceSshdConfigValue GSSAPIAuthentication "GSSAPIAuthentication no"

# Change value X11Forwarding to no.
replaceSshdConfigValue X11Forwarding "X11Forwarding no"

# Remove several tunneling options
replaceSshdConfigValue AllowAgentForwarding "AllowAgentForwarding no"
replaceSshdConfigValue AllowTcpForwarding "AllowTcpForwarding no"
replaceSshdConfigValue PermitTunnel "PermitTunnel no"

# Remove debian banner that gives OS info out
replaceSshdConfigValue DebianBanner "DebianBanner no"

# Change value LogLevel to VERBOSE
replaceSshdConfigValue LogLevel "LogLevel VERBOSE"

# Change value PrintLastLog to yes
replaceSshdConfigValue PrintLastLog "PrintLastLog yes"

# Restart SSH Daemon
sudo systemctl restart sshd
