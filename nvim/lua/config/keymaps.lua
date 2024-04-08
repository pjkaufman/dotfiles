local status_ok, mapkeyper = pcall(require, "util.mapkeyper")
if not status_ok then
  vim.notify("failed to load mapkeyper utils", vim.log.levels.ERROR)
  return
end
local mapkey = mapkeyper.mapkey

-- Silent mapkey option
local opts = { silent = true }

--Remap space as leader key
mapkey("<Space>", "<Nop>", "", opts)

-- Normal --
-- Better window navigation
mapkey("<C-h>", "<C-w>h", "n", opts)
mapkey("<C-j>", "<C-w>j", "n", opts)
mapkey("<C-k>", "<C-w>k", "n", opts)
mapkey("<C-l>", "<C-w>l", "n",opts)

-- Resize with arrows
mapkey("<C-Up>", ":resize -2<CR>", "n", opts)
mapkey("<C-Down>", ":resize +2<CR>", "n", opts)
mapkey("<C-Left>", ":vertical resize -2<CR>", "n", opts)
mapkey("<C-Right>", ":vertical resize +2<CR>", "n", opts)

-- Navigate buffers
mapkey("<S-l>", ":bnext", "n", opts)
mapkey("n", "<S-h>", ":bprevious<CR>", opts)

-- Clear highlights
mapkey("<leader>h", "<cmd>nohlsearch", "n", opts)

-- Close buffers
mapkey("<S-q>", "<cmd>Bdelete!", "n", opts)

-- Insert --
-- Press jk fast to enter
mapkey("jk", "<ESC>", "i", opts)

-- Visual --
-- Stay in indent mode
mapkey("<", "<gv", "v", opts)
mapkey(">", ">gv", "v", opts)
-- Better paste
mapkey("p", '"_dP', "v", opts)
