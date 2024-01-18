install:
	@echo "Building ebook-lint"
	@cargo install --profile release --path . -f

	@echo "Generating bash completion"
	@mkdir -p ${BASH_COMPLETION_USER_DIR}
	@ebook-lint completion bash > "${BASH_COMPLETION_USER_DIR}/ebook-lint-completion"
	@echo "ebook-lint installed" 

build:
	@cargo build
