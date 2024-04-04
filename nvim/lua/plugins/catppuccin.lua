return {
  "catppuccin/nvim",
  name = "catppuccin",
  lazy = false,
  priority = 900,
  config = function()
    local status_ok, _ = pcall(vim.cmd.colorscheme, "catppuccin")
    
    if not status_ok then
      vim.notify("failed to load catppuccin colorscheme", vim.log.levels.ERROR)
      return
    end
  end
}