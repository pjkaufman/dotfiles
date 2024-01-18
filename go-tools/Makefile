
.PHONY: test install cover

test:
	go test ~/dotfiles/go-tools/... -tags "unit"

# this is just meant to give an idea whether or not something has tests in it.
# It is not meant to be used for 100% test coverage. Some folders will be better tested than others.
cover:
	go test -cover ./... -tags "unit"

install:
	@echo "Building go tools"
	@go build -o ~/.local/bin/ebook-lint ./ebook-lint/main.go
	@go build -o ~/.local/bin/git-helper ./git-helper/main.go
	@go build -o ~/.local/bin/song-converter ./song-converter/main.go
	@go build -o ~/.local/bin/cat-ascii ./cat-ascii/main.go
	
	@mkdir -p ${BASH_COMPLETION_USER_DIR}

	@echo "Generating the bash completion for the tools"
	@ebook-lint completion bash > "${BASH_COMPLETION_USER_DIR}/ebook-lint-completion"
	@git-helper completion bash > "${BASH_COMPLETION_USER_DIR}/git-helper-completion"
	@song-converter completion bash > "${BASH_COMPLETION_USER_DIR}/song-converter-completion"
	@cat-ascii completion bash > "${BASH_COMPLETION_USER_DIR}/cat-ascii-completion"
