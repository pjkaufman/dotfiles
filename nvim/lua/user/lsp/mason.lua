local servers = {
	"sumneko_lua",
	"cssls",
	"html",
	"tsserver",
	"pyright",
	"gopls",
	"bashls",
	"jsonls",
	"yamlls",
}

local settings = {
	ui = {
		border = "none",
		icons = {
			package_installed = "◍",
			package_pending = "◍",
			package_uninstalled = "◍",
		},
	},
	log_level = vim.log.levels.INFO,
	max_concurrent_installers = 4,
}

require("mason").setup(settings)
require("mason-lspconfig").setup({
	ensure_installed = servers,
	automatic_installation = true,
})

local lspconfig_status_ok, lspconfig = pcall(require, "lspconfig")
if not lspconfig_status_ok then
	vim.notify("failed to load lspconfig plugin", vim.log.levels.ERROR)
	return
end

local opts = {}

for _, server in pairs(servers) do
	opts = {
		on_attach = require("user.lsp.handlers").on_attach,
		capabilities = require("user.lsp.handlers").capabilities,
		root_dir = function(fname)
			local util = require("lspconfig").util
			return util.root_pattern(".git")(fname)
				or util.root_pattern("tsconfig.base.json")(fname)
				or util.root_pattern("package.json")(fname)
				or util.root_pattern(".eslintrc.js")(fname)
				or util.root_pattern(".eslintrc.json")(fname)
				or util.root_pattern("tsconfig.json")(fname)
				or util.root_pattern("go.mod")(fname)
		end,
	}

	server = vim.split(server, "@")[1]

	local require_ok, conf_opts = pcall(require, "user.lsp.settings." .. server)
	if require_ok then
		opts = vim.tbl_deep_extend("force", conf_opts, opts)
	end

	lspconfig[server].setup(opts)
end
