
.PHONY: test install

completionFile = /tmp/completion

test:
	go test ~/dotfiles/go-tools/... -tags "unit"

install:
	go build -o ~/.local/bin/epub-lint ./epub-lint/main.go
	go build -o ~/.local/bin/git-helper ./git-helper/main.go
	[ -f $(completionFile) ] && rm $(completionFile) 
	epub-lint completion bash > $(completionFile)
	git-helper completion bash >> $(completionFile)
