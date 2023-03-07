local status_ok, whichkey = pcall(require, "which-key")
if not status_ok then
	return
end

local conf = {
	window = {
		border = "single", -- none, single, double, shadow
		position = "bottom", -- bottom, top
	},
}
whichkey.setup(conf)

local n_opts = {
	mode = "n", -- Normal mode
	prefix = "<leader>",
	buffer = nil, -- Global mappings. Specify a buffer number for buffer local mappings
	silent = true, -- use `silent` when creating keymaps
	noremap = true, -- use `noremap` when creating keymaps
	nowait = false, -- use `nowait` when creating keymaps
}

-- local v_opts = {
-- 	mode = "v", -- Visual mode
-- 	prefix = "<leader>",
-- 	buffer = nil, -- Global mappings. Specify a buffer number for buffer local mappings
-- 	silent = true, -- use `silent` when creating keymaps
-- 	noremap = true, -- use `noremap` when creating keymaps
-- 	nowait = false, -- use `nowait` when creating keymaps
-- }

local x_opts = {
	mode = "x", -- Visual Block mode
	prefix = "<leader>",
	buffer = nil, -- Global mappings. Specify a buffer number for buffer local mappings
	silent = true, -- use `silent` when creating keymaps
	noremap = true, -- use `noremap` when creating keymaps
	nowait = false, -- use `nowait` when creating keymaps
}

local function normal_keymap()
	local keymap = {
		["w"] = { "<cmd>update!<CR>", "Save" },
		["q"] = { "<cmd>lua require('utils').quit()<CR>", "Quit" },
		-- ["t"] = { "<cmd>ToggleTerm<CR>", "Terminal" },
		["/"] = { "<cmd>lua require('Comment.api').toggle.linewise.current()<CR>", "Toggle Comment" },
		["h"] = { "<cmd>nohlsearch<CR>", "Clear Highlights" },
		l = {
			name = "LSP",
			f = { "<cmd>lua vim.lsp.buf.format{ async = true }<cr>", "Format" },
		},
		b = {
			name = "Buffer",
			c = { "<Cmd>BDelete this<Cr>", "Close Buffer" },
			f = { "<Cmd>bdelete!<Cr>", "Force Close Buffer" },
			D = { "<Cmd>BWipeout other<Cr>", "Delete All Buffers" },
			b = { "<Cmd>BufferLinePick<Cr>", "Pick a Buffer" },
			p = { "<Cmd>BufferLinePickClose<Cr>", "Pick & Close a Buffer" },
			m = { "<Cmd>JABSOpen<Cr>", "Menu" },
		},
		d = {
			name = "Debug",
			b = { "<cmd>lua require'dap'.toggle_breakpoint()<cr>", "Toggle Breakpoint" },
			c = { "<cmd>lua require'dap'.continue()<cr>", "Continue" },
			i = { "<cmd>lua require'dap'.step_into()<cr>", "Step Into" },
			o = { "<cmd>lua require'dap'.step_over()<cr>", "Step Over" },
			O = { "<cmd>lua require'dap'.step_out()<cr>", "Step Out" },
			r = { "<cmd>lua require'dap'.repl.toggle()<cr>", "Repl Toggle" },
			l = { "<cmd>lua require'dap'.run_last()<cr>", "Run Last" },
			u = { "<cmd>lua require'dapui'.toggle()<cr>", "Dap UI Toggle" },
			q = { "<cmd>lua require'dap'.terminate()<cr>", "Quit" },
			t = { "<cmd>lua require'dap-go'.debug_test()<cr>", "Go Test" },
		},
		f = {
			name = "Find",
			f = { ":Telescope find_files<CR>", "Files" },
			t = { ":Telescope live_grep<CR>", "Text" },
			p = { ":Telescope projects<CR>", "Projects" },
			b = { ":Telescope buffers<CR>", "Buffers" },
			e = { "<cmd>NvimTreeToggle<cr>", "Explorer" },
		},
		z = {
			name = "System",
			i = { "<cmd>PackerInstall<cr>", "Install" },
			p = { "<cmd>PackerProfile<cr>", "Profile" },
			s = { "<cmd>PackerSync<cr>", "Sync" },
			S = { "<cmd>PackerStatus<cr>", "Status" },
			u = { "<cmd>PackerUpdate<cr>", "Update" },
		},
		g = {
			name = "Git",
			g = { "<cmd>lua _LAZYGIT_TOGGLE()<CR>", "Toggle Lazy Git" },
		},
	}

	whichkey.register(keymap, n_opts)
end

-- local function visual_keymap()
-- 	local keymap = {
-- 		r = {
-- 			name = "Refactor",
-- 			f = { [[<cmd>lua require('refactoring').refactor('Extract Function')<cr>]], "Extract Function" },
-- 			F = {
-- 				[[ <cmd>lua require('refactoring').refactor('Extract Function to File')<cr>]],
-- 				"Extract Function to File",
-- 			},
-- 			v = { [[<cmd>lua require('refactoring').refactor('Extract Variable')<cr>]], "Extract Variable" },
-- 			i = { [[<cmd>lua require('refactoring').refactor('Inline Variable')<cr>]], "Inline Variable" },
-- 			r = { [[<cmd>lua require('telescope').extensions.refactoring.refactors()<cr>]], "Refactor" },
-- 			d = { [[<cmd>lua require('refactoring').debug.print_var({})<cr>]], "Debug Print Var" },
-- 		},
-- 	}
--
-- 	whichkey.register(keymap, v_opts)
-- end

local function visual_block_keymap()
	local keymap = {
		["/"] = { "<esc><cmd>lua require('Comment.api').toggle.linewise(vim.fn.visualmode())<CR>", "Toggle Comment" },
	}

	whichkey.register(keymap, x_opts)
end

-- local function code_keymap()
-- 	vim.api.nvim_create_autocmd("FileType", {
-- 		pattern = "*",
-- 		callback = function()
-- 			vim.schedule(CodeRunner)
-- 		end,
-- 	})
--
-- 	function CodeRunner()
-- 		local bufnr = vim.api.nvim_get_current_buf()
-- 		local ft = vim.api.nvim_buf_get_option(bufnr, "filetype")
-- 		local fname = vim.fn.expand("%:p:t")
-- 		local keymap_c = {} -- normal key map
-- 		local keymap_c_v = {} -- visual key map
--
-- 		if ft == "python" then
-- 			keymap_c = {
-- 				name = "Code",
-- 				-- r = { "<cmd>update<CR><cmd>exec '!python3' shellescape(@%, 1)<cr>", "Run" },
-- 				-- r = { "<cmd>update<CR><cmd>TermExec cmd='python3 %'<cr>", "Run" },
-- 				i = { "<cmd>cexpr system('refurb --quiet ' . shellescape(expand('%'))) | copen<cr>", "Inspect" },
-- 				r = {
-- 					"<cmd>update<cr><cmd>lua require('utils.term').open_term([[python3 ]] .. vim.fn.shellescape(vim.fn.getreg('%'), 1), {direction = 'float'})<cr>",
-- 					"Run",
-- 				},
-- 				m = { "<cmd>TermExec cmd='nodemon -e py %'<cr>", "Monitor" },
-- 			}
-- 		elseif ft == "lua" then
-- 			keymap_c = {
-- 				name = "Code",
-- 				r = { "<cmd>luafile %<cr>", "Run" },
-- 			}
-- 		elseif ft == "rust" then
-- 			keymap_c = {
-- 				name = "Code",
-- 				r = { "<cmd>execute 'Cargo run' | startinsert<cr>", "Run" },
-- 				D = { "<cmd>RustDebuggables<cr>", "Debuggables" },
-- 				h = { "<cmd>RustHoverActions<cr>", "Hover Actions" },
-- 				R = { "<cmd>RustRunnables<cr>", "Runnables" },
-- 			}
-- 		elseif ft == "go" then
-- 			keymap_c = {
-- 				name = "Code",
-- 				r = { "<cmd>GoRun<cr>", "Run" },
-- 			}
-- 		elseif ft == "typescript" or ft == "typescriptreact" or ft == "javascript" or ft == "javascriptreact" then
-- 			keymap_c = {
-- 				name = "Code",
-- 				o = { "<cmd>TypescriptOrganizeImports<cr>", "Organize Imports" },
-- 				r = { "<cmd>TypescriptRenameFile<cr>", "Rename File" },
-- 				i = { "<cmd>TypescriptAddMissingImports<cr>", "Import Missing" },
-- 				F = { "<cmd>TypescriptFixAll<cr>", "Fix All" },
-- 				u = { "<cmd>TypescriptRemoveUnused<cr>", "Remove Unused" },
-- 				R = { "<cmd>lua require('config.test').javascript_runner()<cr>", "Choose Test Runner" },
-- 				-- s = { "<cmd>2TermExec cmd='yarn start'<cr>", "Yarn Start" },
-- 				-- t = { "<cmd>2TermExec cmd='yarn test'<cr>", "Yarn Test" },
-- 			}
-- 		elseif ft == "java" then
-- 			keymap_c = {
-- 				name = "Code",
-- 				o = { "<cmd>lua require'jdtls'.organize_imports()<cr>", "Organize Imports" },
-- 				v = { "<cmd>lua require('jdtls').extract_variable()<cr>", "Extract Variable" },
-- 				c = { "<cmd>lua require('jdtls').extract_constant()<cr>", "Extract Constant" },
-- 				t = { "<cmd>lua require('jdtls').test_class()<cr>", "Test Class" },
-- 				n = { "<cmd>lua require('jdtls').test_nearest_method()<cr>", "Test Nearest Method" },
-- 			}
-- 			keymap_c_v = {
-- 				name = "Code",
-- 				v = { "<cmd>lua require('jdtls').extract_variable(true)<cr>", "Extract Variable" },
-- 				c = { "<cmd>lua require('jdtls').extract_constant(true)<cr>", "Extract Constant" },
-- 				m = { "<cmd>lua require('jdtls').extract_method(true)<cr>", "Extract Method" },
-- 			}
-- 		end
--
-- 		if fname == "package.json" then
-- 			keymap_c.v = { "<cmd>lua require('package-info').show()<cr>", "Show Version" }
-- 			keymap_c.c = { "<cmd>lua require('package-info').change_version()<cr>", "Change Version" }
-- 			-- keymap_c.s = { "<cmd>2TermExec cmd='yarn start'<cr>", "Yarn Start" }
-- 			-- keymap_c.t = { "<cmd>2TermExec cmd='yarn test'<cr>", "Yarn Test" }
-- 		end
--
-- 		if fname == "Cargo.toml" then
-- 			keymap_c.u = { "<cmd>lua require('crates').upgrade_all_crates()<cr>", "Upgrade All Crates" }
-- 		end
--
-- 		if next(keymap_c) ~= nil then
-- 			local k = { c = keymap_c }
-- 			local o = { mode = "n", silent = true, noremap = true, buffer = bufnr, prefix = "<leader>", nowait = true }
-- 			whichkey.register(k, o)
-- 			-- legendary.bind_whichkey(k, o, false)
-- 		end
--
-- 		if next(keymap_c_v) ~= nil then
-- 			local k = { c = keymap_c_v }
-- 			local o = { mode = "v", silent = true, noremap = true, buffer = bufnr, prefix = "<leader>", nowait = true }
-- 			whichkey.register(k, o)
-- 			-- legendary.bind_whichkey(k, o, false)
-- 		end
-- 	end
-- end

normal_keymap()
-- visual_keymap()
visual_block_keymap()
-- code_keymap()
