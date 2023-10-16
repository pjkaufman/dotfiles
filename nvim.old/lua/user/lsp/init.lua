local status_ok, _ = pcall(require, "lspconfig")
if not status_ok then
  vim.notify("failed to load lspconfig plugin", vim.log.levels.ERROR)
  return
end

require("user.lsp.mason")
require("user.lsp.handlers").setup()
require("user.lsp.null-ls")
