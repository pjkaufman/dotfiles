-- https://github.com/neovim/nvim-lspconfig/blob/master/lua/lspconfig/server_configurations/gopls.lua
return {
	cmd = { "gopls" },
	settings = {
		gopls = {
			buildFlags = { "-tags=unit integration" },
			usePlaceholders = false,
			completeUnimported = true,
			analyses = {
				nilness = true,
				shadow = true,
				unusedparams = true,
				unusewrites = true,
			},
			staticcheck = true,
			codelenses = {
				generate = false,
				gc_details = true,
				test = true,
				tidy = true,
			},
		},
	},
}
