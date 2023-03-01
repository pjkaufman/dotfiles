#!/bin/bash

# prepend_text_to_file takes the specified text and adds it to the start of the file specified
# $1 is the text to add to the start of the file
# $2 is the path for the file to put the text at the start of
prepend_text_to_file() {
  echo "$1" > /tmp/tmpfile.$$
  cat "$2" >> /tmp/tmpfile.$$
  mv /tmp/tmpfile.$$ "$2"
}

# make sure that if the computer type is missing we get the type from the user exporting it so we can use it throughout the install script
# while we just print out the type if we already have it

if [ -z "${COMPUTER_TYPE}" ]
then
  read -p 'Is this a personal computer? [y]es or [n]o: ' response_char

  if [ response_char = "y" ]
  then
    prepend_text_to_file 'export COMPUTER_TYPE=personal' ~/.local_extra
    export COMPUTER_TYPE=personal
  else
    prepend_text_to_file 'export COMPUTER_TYPE=work' ~/.local_extra
    export COMPUTER_TYPE=work
  fi
else
  echo "The computer is a ${COMPUTER_TYPE} one."
fi
