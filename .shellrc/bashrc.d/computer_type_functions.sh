#!/bin/bash

# return 0 or 1 based on https://unix.stackexchange.com/a/348132

TRUE=0
FALSE=1

is_work_computer() {
  if [ ${COMPUTER_TYPE} = "work" ]; then
    return $TRUE
  fi

  return $FALSE
}

is_personal_computer() {
  if [ ${COMPUTER_TYPE} = "work" ]; then
    return $FALSE
  fi

  return $TRUE
}
