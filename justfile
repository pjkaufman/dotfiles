# shellcheck disable=SC2148

alias i := install

install:
  ./install.sh
format:
  shfmt -w .
lint:
  # shellcheck disable=SC2046
  shellcheck -x ./install.sh $(find ./bin/* -maxdepth 1 -not -name '*.md';) ./bash/bashrc ./bash/functions/*.sh ./install/*.sh
  shfmt -d .