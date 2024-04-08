return {
  "nvim-treesitter/nvim-treesitter",
  build = ":TSUpdate",
  config = function()
    local status_ok, treesitter = pcall(require, "nvim-treesitter")
    if not status_ok then
      vim.notify("failed to load nvim-treesitter plugin", vim.log.levels.ERROR)
      return
    end

    local status_ok, configs = pcall(require, "nvim-treesitter.configs")
    if not status_ok then
      vim.notify("failed to load nvim-treesitter config", vim.log.levels.ERROR)
      return
    end

    configs.setup({
      ensure_installed = {
        "lua",
        "markdown",
        "markdown_inline",
        "bash",
        "python",
        "go",
        "javascript",
        "proto",
        "sql",
        "typescript",
        "jsdoc",
        "comment",
        "html",
        "java",
        "json",
        "make",
        "yaml",
      }, -- put the language you want in this array
      -- ensure_installed = "all", -- one of "all" or a list of languages
      ignore_install = { "" }, -- List of parsers to ignore installing
      sync_install = false, -- install languages synchronously (only applied to `ensure_installed`)
      highlight = {
        enable = true, -- false will disable the whole extension
        disable = { "css" }, -- list of language that will be disabled
      },
      autopairs = {
        enable = true,
      },
      indent = { enable = true, disable = { "python", "css" } },
      context_commentstring = {
        enable = true,
        enable_autocmd = false,
      },
    })
  end
}