# Dotfiles

This is a repo which contains the config files and several scripts I tend to use to help me in my linux environments.

## Installation

When installing things, you should just be able to run `install.sh`. It should do the trick for installing all of the needed applications for the specified environment type.
You may need to run `chmod +x` on `install.sh` in order to run the script.

If this is a personal computer and it does not contain flatpak, you will need to manually setup flatpak and then run the installation again.

_Note that the expectation is that this repo will be in the home directory of the current user. If it is not, the installation may fail._

### Tmux

Once you start up tmux for the first time, make sure to enter `ctrl+a+I` in a tmux session to install the tmux plugins.

## Rational

It can be hard and time consuming to setup one's applications and environment across multiple computers.
This script and these configs allow me to install the base level of the environment which can then be tinkered with from there.

The repo also acts as a secondary copy of my configs which allows me to backup my configs. If something were to happen to my computer, this allows for an easier setup for new environments.

## Dependencies

These dotfiles are meant to be run in bash. As such they are meant to be run on a unix system and not Windows.

### Programs

The current list of programs that need installing and are used are as follows:

| Program Name | Installation Method | Use Case |
| ------------ | ------------------- | -------- |
| `grep` | apt | General cli utility for string searching |
| `ripgrep` | apt | General cli utility for string searching that is used in some NeoVim plugins |
| `curl` | apt | General cli utility for getting webpage content from the cli |
| `fzf` | apt | Fuzzy finder for the cli |
| `btop` | apt | Cli task/resource manager |
| `python3` | apt | Helps with writing some scripts and installing some programs used |
| `pip3` | apt | Helps get some packages that are not available on in apt and are written in Python |
| `rename` | apt | Helps rename files using regex from the cli |


## Known Issues

### Neovim

- The clipboard does not seem to connect to the system clipboard properly
- Running Go tests does not seem to work
- golangci-lint does not seem to work for Go


## Setting Up Calibre DeDRM

- Install Calibre
- Install a compatible version of the DeDRM plugin
- Install Kindle for PC version 1.17 via wine
- Install Calibre in wine*
- Install a compatible version of the DeDRM plugin in wine*
- Restart Calibre in wine*
  - The DeDRM plugin should autoload the key for DeDRM the books
  - If it does not, press the green plus sign and it should load the key
- Export the key
- Import the key in the version of Caliber on your non-wine install
- Load any books you want from Kindle for PC into Calibre
- They should now be DRM free

*: the wine version is needed for working with Kindle for PC in wine if your non-wine Calibre cannot detect the key for Kindle for PC 

_Note: this method is only meant to be used for books you have bought. This is not meant to be something used on books you have not bought._
