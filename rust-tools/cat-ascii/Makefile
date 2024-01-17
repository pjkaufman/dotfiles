install:
	@echo "Building cat-ascii"
	@cargo install --profile release --path . -f

	@echo "Generating bash completion"
	@mkdir -p ${BASH_COMPLETION_USER_DIR}
	@cat-ascii completion bash > "${BASH_COMPLETION_USER_DIR}/cat-ascii-completion"
	@echo "cat-ascii installed" 

build:
	@cargo build
