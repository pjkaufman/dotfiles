local status_ok, keymapper = pcall(require, "util.keymapper")
if not status_ok then
  vim.notify("failed to load keymapper utils", vim.log.levels.ERROR)
  return
end
local mapkey = keymapper.mapkey

return {
  "nvim-telescope/telescope.nvim",
  lazy = false,
  dependencies = { "nvim-lua/plenary.nvim" },
  config = function()
    local status_ok, telescope = pcall(require, "telescope")
    if not status_ok then
      vim.notify("failed to load telescope plugin", vim.log.levels.ERROR)
      return
    end

    local actions = require("telescope.actions")

    telescope.setup({
      defaults = {

        prompt_prefix = " ",
        selection_caret = " ",
        path_display = { "smart" },
        file_ignore_patterns = { ".git/", "node_modules" },

        mappings = {
          i = {
            ["<Down>"] = "move_selection_next",
            ["<Up>"] = "move_selection_previous"
            -- ["<C-j>"] = "move_selection_next",
            -- ["<C-k>"] = "move_selection_previous",
          },
        },
      },
    })
  end,
  keys = {
    -- Make sure that the Nvim Tree is closed before we opene find files for telescope which lets finding files when none is currently open work instead of just closing neovim
		-- from https://www.reddit.com/r/neovim/comments/v49jxi/dont_open_telescope_action_in_nvimtree_window/
    mapkey("<leader>ff", "execute 'NvimTreeClose' | Telescope find_files", "n"),
    mapkey("<leader>ft", "Telescope live_grep", "n"),
    mapkey("<leader>fp", "Telescope projects", "n"),
    mapkey("<leader>fb", "Telescope buffers", "n"),
  }
}