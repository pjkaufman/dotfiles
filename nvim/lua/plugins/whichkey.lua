return {
  "folke/which-key.nvim",
  lazy = false,
  config = function()
    local status_ok, whichkey = pcall(require, "which-key")
    if not status_ok then
      vim.notify("failed to load which-key plugin", vim.log.levels.ERROR)
      return
    end
    
    local conf = {
      window = {
        border = "single", -- none, single, double, shadow
        position = "bottom", -- bottom, top
      },
    }
    whichkey.setup(conf)
  end
}