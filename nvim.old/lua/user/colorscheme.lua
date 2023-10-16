local colorscheme = "catppuccin"

local status_ok, _ = pcall(vim.cmd.colorscheme, colorscheme)
if not status_ok then
	vim.notify("failed to load catppuccin colorscheme", vim.log.levels.ERROR)
	return
end
