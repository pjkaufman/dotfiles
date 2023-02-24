# Load the shell dotfiles, and then some:
# * ~/.path can be used to extend `$PATH`.
# * ~/.extra_exports can be used for other settings you don’t want to commit.
for file in ~/.{extra_exports,functions,path,exports}; do
  [ -r "$file" ] && [ -f "$file" ] && . "$file";
done;
unset file;

# if running bash
if [ -n "$BASH_VERSION" ]; then
  # include .bashrc if it exists
  if [ -f "$HOME/.bashrc" ]; then
	  . "$HOME/.bashrc"
  fi
fi

# Case-insensitive globbing (used in pathname expansion)
shopt -s nocaseglob;

# Append to the Bash history file, rather than overwriting it
shopt -s histappend;

# Autocorrect typos in path names when using `cd`
shopt -s cdspell;

# Enable some Bash 4 features when possible:
# * `autocd`, e.g. `**/qux` will enter `./foo/bar/baz/qux`
# * Recursive globbing, e.g. `echo **/*.txt`
for option in autocd globstar; do
	shopt -s "$option" 2> /dev/null;
done;

#SSH for Github
eval $(ssh-agent -s)

# gvm use go1.18
# set nvm version to use if it does not exist already
if [ -d "$HOME/.nvm" ] ; then
  nvm use 16.0
fi

# ~/.extra can be used for other settings you don’t want to commit.
[ -r "$HOME/.extra" ] && [ -f "$HOME/.extra" ] && . "$HOME/.extra";

# clear the terminal after the setup is over
c
