# Dotfiles

This is a repo which contains the config files and several scripts I tend to use to help me in my linux environments.

## Installation

When installing things, you should just be able to run `install.sh`. It should do the trick for installing all of the needed applications for the specified environment type.
You may need to run `chmod +x` on `install.sh` in order to run the script.

If this is a personal computer and it does not contain flatpak, you will need to manually setup flatpak and then run the installation again.

_Note that the expectation is that this repo will be in the home directory of the current user. If it is not, the installation may fail._

## Rational

It can be hard and time consuming to setup one's applications and environment across multiple computers.
This script and these configs allow me to install the base level of the environment which can then be tinkered with from there.

The repo also acts as a secondary copy of my configs which allows me to backup my configs. If something were to happen to my computer, this allows for an easier setup for new environments.

## Dependencies

## Known Issues

### Neovim

- The clipboard does not seem to connect to the system clipboard properly
- Running Go tests does not seem to work
- golangci-lint does not seem to work for Go

### Setting Up Calibre DeDRM

First you will need to install Calibre and a compatible version of the DeDRM plugin. Make sure to install Kindle for PC version 1.17 via wine. You will also want to install a calibre version along with its compatible version of DeDRM into wine and then restart Calibre once the plugin is present. From there it should autoload the key and if not press the green plus sign and it should load the key. Once that is done, export the key and then import that key in the version of calibre on your non-wine portion. Then go ahead and import any books you would like from Kindle for PC. They should all be DRM free now.
