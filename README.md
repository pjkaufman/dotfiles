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
| `grep` | dnf | General cli utility for string searching |
| `curl` | dnf | General cli utility for getting webpage content from the cli |
| `fzf` | dnf | Fuzzy finder for the cli |
| `btop` | dnf | Cli task/resource manager |
| `python3` | dnf | Helps with writing some scripts and installing some programs used |
| `pip3` | dnf | Helps get some packages that are not available on in dnf and are written in Python |
| `rename` | dnf | Helps rename files using regex from the cli |
| `bat` | dnf | Adds syntax highlighting to cat commands |
| `imgp` | dnf | A nice cli image compressor that I keep around for the time being |
| `evince` | dnf | A nice pdf viewer that can be launched from the cli |
| `Brave` | flatpak | A chromium based browser |
| `Minecraft` | flatpak | Minecraft game |
| `GnuCash` | flatpak | A local way to do budgeting and track expenses |
| `Calibre` | flatpak | An all in one ebook editor, viewer, and library |
| `Obsidian` | flatpak | A good local first not taking app that is very extendable and a scratchpad for Spanish writing |
| `git` | dnf | Cli program for interacting with git repos |
| `openssh-server` | dnf | Program for doing ssh key generation use for git authentication |
| `go` | script | Golang is a great program for developing scripts and programs that are more performant |
| `light` | dnf | Program for handing brightness changes |
| `pulseaudio-utils` | dnf | Program for updating sound settings |
| `kitty` | dnf | Default terminal |
| `goimports` | Golang | A Golang file modifier that adds missing imports where possible |
| `dlv` | Golang | A Golang debugger server |
| `eslint` | NPM | A JS/TS file formatter/linter |
| `stylua` | Cargo | A Lua file formatter/linter |
| `nvm` | script | A file for managing node versions |
| `rkhunter` | dnf | A virus scanner and security checker for Linux |
| `syncthing` | dnf | Program for syncing local files across devices more easily than would otherwise be possible |
| `tmux` | dnf | Program for easier session management in the cli |
| `tpm` | script | Plugin manager for tmux |
| `weasyprint` | dnf | Convert html to pdf |
| `pandoc` | dnf | Document converter |

## Known Issues

- VsCode is always updating a config when it launches

## TODOs

- Add option to convert church songs to odf or similar format from html for flexibility
- Add Obsidian configs if possible
