local status_ok, keymapper = pcall(require, "util.keymapper")
if not status_ok then
  vim.notify("failed to load keymapper utils", vim.log.levels.ERROR)
  return
end
local mapkey = keymapper.mapkey

return {
  "nvim-tree/nvim-tree.lua",
	lazy = false,
  config = {},
  keys = {
    mapkey("<leader>fe", "NvimTreeToggle", "n"),
  },
}