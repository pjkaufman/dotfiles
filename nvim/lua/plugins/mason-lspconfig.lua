local opts = {
	ensure_installed = {
		"lua_ls",
    "cssls",
    "html",
    "tsserver",
    "pyright",
    "gopls",
    "bashls",
    "jsonls",
    "yamlls",
	},

	automatic_installation = true,
}

return {
	"williamboman/mason-lspconfig.nvim",
	opts = opts,
	event = "BufReadPre",
	dependencies = "williamboman/mason.nvim",
}