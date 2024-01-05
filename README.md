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
| `imgp` | apt | A nice cli image compressor used in my epub compression scripts |
| `flameshot` | apt* | A useful screenshot utility |
| `evince` | apt | A nice pdf viewer that can be launched from the cli |
| `Brave` | flatpak | A chromium based browser |
| `Minecraft` | flatpak | Minecraft game |
| `GnuCash` | flatpak | A local way to do budgeting and track expenses |
| `Calibre` | flatpak | An all in one ebook editor, viewer, and library |
| `Obsidian` | flatpak | A good local first not taking app that is very extendable and a scratchpad for Spanish writing |
| `git` | PPA | Cli program for interacting with git repos |
| `openssh-server` | apt | Program for doing ssh key generation use for git authentication |
| `go` | script | Golang is a great program for developing scripts and programs that are more performant |
| `i3` | apt | Window manager for X11 |
| `scrot` | apt | Program for taking screenshots in X11 |
| `light` | apt | Program for handing brightness changes |
| `feh` | apt | Program for setting the background image on X11 |
| `i3lock` | apt | Program for setting the lock screen on X11 for i3 |
| `i3status` | apt | Program helps with setting status info in X11 for i3 |
| `dmenu` | apt | Program for X11 that helps with selecting apps | 
| `pulseaudio-utils` | apt | Program for updating sound settings |
| `bumblebee-status` | pip3 | Program for setting a better status bar on X11 for i3 |
| `kitty` | apt | Default terminal |
| `NeoVim` | PPA | A personal development editor for editing different kinds of files |
| `codespell` | pip3 | A linter for spelling correction for NeoVim |
| `black` | pip3 | A linter for python used in NeoVim |
| `flake8` | pip3 | A python formatter used in NeoVim |
| `beautysh` | pip3 | A bash formatter for bash/sh files in NeoVim |
| `protolint` | Golang | A proto file formatter/linter for NeoVim |
| `golangci-lint` | Golang | A Golang file linter for NeoVim |
| `goimports` | Golang | A Golang file modifier that adds missing imports where possible |
| `dlv` | Golang | A Golang debugger server |
| `eslint` | NPM | A JS/TS file formatter/linter |
| `stylua` | Cargo | A Lua file formatter/linter |
| `nvm` | script | A file for managing node versions |
| `rkhunter` | apt | A virus scanner and security checker for Linux |
| `cargo` | apt | A Rust package manager used to install a couple of programs |
| `topgrade` | Cargo | A package that helps manage updating many package managers |
| `cargo-update` | Cargo | A package for helping update Rust and Rust packages |
| `syncthing` | PPA | Program for syncing local files across devices more easily than would otherwise be possible |
| `tmux` | apt | Program for easier session management in the cli |
| `tpm` | script | Plugin manager for tmux |
| `wine` | PPA | Windows Program Emulator used for Kindle for PC since no Linux version is available |
| `wkhtmltopdf` | script | Program for converting html files to a pdf with an autogenerated table of contents |
| `openssl`, `libssl-dev`, `libxrender-dev`, `libx11-dev`, `libxext-dev`, `libfontconfig1-dev`, `libfreetype6-dev`, `fontconfig` | apt | Dependencies for `wkhtmltopdf` |
| `pandoc` | apt | Document converter |

*: This is currently not really working on wayland

## Known Issues

### Neovim

- The clipboard does not seem to connect to the system clipboard properly
- Running Go tests does not seem to work
- golangci-lint does not seem to work for Go

## TODOs

- Add VsCode
- Convert NeoVim to a more stable setup
- Setup monitor in Sway
- Add option to convert church songs to odf or similar format from html for flexibility
- Add Obsidian configs if possible
