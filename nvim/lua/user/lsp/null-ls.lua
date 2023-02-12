local null_ls_status_ok, null_ls = pcall(require, "null-ls")
if not null_ls_status_ok then
	return
end

-- https://github.com/jose-elias-alvarez/null-ls.nvim/tree/main/lua/null-ls/builtins/formatting
local formatting = null_ls.builtins.formatting
-- https://github.com/jose-elias-alvarez/null-ls.nvim/tree/main/lua/null-ls/builtins/diagnostics
local diagnostics = null_ls.builtins.diagnostics

local sources = {
	-- formatting
	formatting.black.with({ extra_args = { "--fast" } }),
	formatting.stylua,
	formatting.google_java_format,
	formatting.eslint.with({
		args = {
			"--fix",
			"$FILENAME",
		},
	}),
	formatting.codespell,
	formatting.gofmt,
	formatting.goimports,
	formatting.protolint.with({ extra_args = { "--fix" } }),
	-- diagnostics
	diagnostics.flake8,
	diagnostics.eslint.with({
		method = null_ls.methods.DIAGNOSTICS_ON_SAVE,
	}),
	diagnostics.codespell,
	-- diagnostics.golangci_lint,
	--diagnostics.golangci_lint.with { extra_args = { "--config=.golangci.yaml" } },
	diagnostics.golangci_lint.with({
		args = {
			"run",
			"--enable-all",
			"--disable exhaustivestruct",
			"--out-format json",
			"$DIRNAME",
			"--path-prefix",
			"$ROOT",
		},
	}),
	--diagnostics.golangci_lint.with { { extra_args = { "--config", vim.fn.expand("~/mono/.ci/golangci.yml")} } },
}

local augroup = vim.api.nvim_create_augroup("LspFormatting", {})

-- https://github.com/prettier-solidity/prettier-plugin-solidity
null_ls.setup({
	debug = false,
	sources = sources,
	update_in_insert = true,
	on_attach = function(client, bufnr)
		if client.supports_method("textDocument/formatting") then
			vim.api.nvim_clear_autocmds({ group = augroup, buffer = bufnr })
			vim.api.nvim_create_autocmd("BufWritePre", {
				group = augroup,
				buffer = bufnr,
				callback = function()
					-- on 0.8, you should use vim.lsp.buf.format({ bufnr = bufnr }) instead
					vim.lsp.buf.format({ bufnr = bufnr })
				end,
			})
		end
	end,
})
