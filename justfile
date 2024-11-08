format:
  shfmt -w .
lint:
  shellcheck -x ./install.sh $(find ./bin/* -maxdepth 1 -not -name '*.md';) ./bash/bashrc ./bash/functions/*.sh ./install/*.sh
  shfmt -d .