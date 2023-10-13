
.PHONY: test install cover

completionFile = /tmp/completion

test:
	go test ~/dotfiles/go-tools/... -tags "unit"

# this is just meant to give an idea whether or not something has tests in it.
# It is not meant to be used for 100% test coverage. Some folders will be better tested than others.
cover:
	go test -cover ./... -tags "unit"

install:
	go build -o ~/.local/bin/ebook-lint ./ebook-lint/main.go
	go build -o ~/.local/bin/git-helper ./git-helper/main.go
	go build -o ~/.local/bin/song-converter ./song-converter/main.go
	
	ebook-lint completion bash > $(completionFile)
	git-helper completion bash >> $(completionFile)
	song-converter completion bash >> $(completionFile)
