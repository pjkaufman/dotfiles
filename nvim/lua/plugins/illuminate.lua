return {
	"RRethy/vim-illuminate",
	lazy = false,
	config = function()
    local status_ok, illuminate = pcall(require, "illuminate")
    if not status_ok then
      vim.notify("failed to load illuminate plugin", vim.log.levels.ERROR)
      return
    end
		illuminate.configure({})
	end,
}