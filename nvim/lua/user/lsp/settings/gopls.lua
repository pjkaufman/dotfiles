-- https://github.com/neovim/nvim-lspconfig/blob/master/lua/lspconfig/server_configurations/gopls.lua
return {
	cmd = { "gopls"},
	settings = {
		gopls = {
      buildFlags = { "-tags=unit integration" },
      usePlaceholders = false,
			analyses = {
				nilness = true,
				shadow = true,
				unusedparams = true,
				unusewrites = true,
			},
      staticcheck = true,
		},
	},
}
