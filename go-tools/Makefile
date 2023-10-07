
.PHONY: test install

completionFile = /tmp/completion

test:
	go test ~/dotfiles/go-tools/... -tags "unit"

install:
	go build -o ~/.local/bin/epub-lint ./epub-lint/main.go
	go build -o ~/.local/bin/git-helper ./git-helper/main.go
	go build -o ~/.local/bin/song-converter ./song-converter/main.go
	
	epub-lint completion bash > $(completionFile)
	git-helper completion bash >> $(completionFile)
	song-converter completion bash >> $(completionFile)
