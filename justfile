# shellcheck disable=SC2148

alias i := install

install:
  ./install.sh
format:
  shfmt -w .
lint:
  # shellcheck disable=SC2046
  shellcheck -x ./install.sh $(find ./bash -type f -exec file --mime-type {} \; | grep "text/x-shellscript" | cut -d: -f1) $(find ./install -type f -exec file --mime-type {} \; | grep "text/x-shellscript" | cut -d: -f1) $(find ./bin -type f -exec file --mime-type {} \; | grep "text/x-shellscript" | cut -d: -f1)
  shfmt -d .