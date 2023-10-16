local dap_status_ok, dap = pcall(require, "dap")
if not dap_status_ok then
	vim.notify("failed to load dap plugin", vim.log.levels.ERROR)
	return
end

local dap_ui_status_ok, dapui = pcall(require, "dapui")
if not dap_ui_status_ok then
	vim.notify("failed to load dapui plugin", vim.log.levels.ERROR)
	return
end

-- local dap_install_status_ok, dap_install = pcall(require, "dap-install")
-- if not dap_install_status_ok then
-- return
-- end

local dap_go_status_ok, dap_go = pcall(require, "dap-go")
if not dap_go_status_ok then
	vim.notify("failed to load dap-go plugin", vim.log.levels.ERROR)
	return
end

local dap_vt_status_ok, dap_vt = pcall(require, "nvim-dap-virtual-text")
if not dap_vt_status_ok then
	vim.notify("failed to load nvim-dap-virtual-text plugin", vim.log.levels.ERROR)
	return
end

local dap_vscode_status_ok, dap_vscode = pcall(require, "dap-vscode-js")
if not dap_vscode_status_ok then
	vim.notify("failed to load dap-vscode-js plugin", vim.log.levels.ERROR)
	return
end

-- dap_install.setup({})

dap.set_log_level("TRACE")
dap_go.setup()

local function getCWDGolang(_path, _service_name)
	local _, index_end = string.find(string.lower(_path), string.lower(_service_name))
	if index_end == nil then
		vim.notify(
			string.format("failed to get Go cwd since '%s' not in current file path", _service_name),
			vim.log.levels.ERROR
		)

		return _path
	end

	return string.sub(_path, 0, index_end)
end

table.insert(dap.configurations["go"], {
	type = "go",
	name = "Debug Microservice",
	request = "launch",
	program = function()
		local service_name = vim.fn.input("Service Name: ")
		local file_path = vim.api.nvim_buf_get_name(0)
		local cwd = getCWDGolang(file_path, service_name)
		if file_path == cwd then
			return "${file}"
		end

		local executable_name = vim.fn.input("Executable Name: ")
		local program_folder = string.format("%s/cmd/%s", cwd, executable_name)

		return program_folder
	end,
	cwd = function()
		local service_name = vim.fn.input("Service Name: ")
		local file_path = vim.api.nvim_buf_get_name(0)
		local cwd = getCWDGolang(file_path, service_name)

		return cwd
	end,
})

dap_vt.setup()

-- dap_install.config("python", {})
-- add other configs here

dapui.setup({
	expand_lines = true,
	icons = { expanded = "", collapsed = "", circular = "" },
	mappings = {
		-- Use a table to apply multiple mappings
		expand = { "<CR>", "<2-LeftMouse>" },
		open = "o",
		remove = "d",
		edit = "e",
		repl = "r",
		toggle = "t",
	},
	layouts = {
		{
			elements = {
				{ id = "scopes", size = 0.33 },
				{ id = "breakpoints", size = 0.17 },
				{ id = "stacks", size = 0.25 },
				{ id = "watches", size = 0.25 },
			},
			size = 0.33,
			position = "right",
		},
		{
			elements = {
				{ id = "repl", size = 0.45 },
				{ id = "console", size = 0.55 },
			},
			size = 0.27,
			position = "bottom",
		},
	},
	floating = {
		max_height = 0.9,
		max_width = 0.5, -- Floats will be treated as percentage of your screen.
		border = vim.g.border_chars, -- Border style. Can be 'single', 'double' or 'rounded'
		mappings = {
			close = { "q", "<Esc>" },
		},
	},
})

vim.fn.sign_define("DapBreakpoint", { text = "", texthl = "DiagnosticSignError", linehl = "", numhl = "" })

dap.listeners.after.event_initialized["dapui_config"] = function()
	dapui.open()
end

dap.listeners.before.event_terminated["dapui_config"] = function()
	dapui.close()
end

dap.listeners.before.event_exited["dapui_config"] = function()
	dapui.close()
end

dap_vscode.setup({
	adapters = { "pwa-node", "pwa-chrome", "pwa-msedge", "node-terminal", "pwa-extensionHost" },
})

for _, language in ipairs({ "typescript", "javascript" }) do
	dap.configurations[language] = {
		{
			type = "pwa-node",
			request = "launch",
			name = "Launch file",
			program = "${file}",
			cwd = "${workspaceFolder}",
		},
		{
			type = "pwa-node",
			request = "attach",
			name = "Attach",
			processId = require("dap.utils").pick_process,
			cwd = "${workspaceFolder}",
		},
		{
			type = "pwa-node",
			request = "launch",
			name = "Debug Jest Tests",
			-- trace = true, -- include debugger info
			runtimeExecutable = "node",
			runtimeArgs = {
				"./node_modules/jest/bin/jest.js",
				"--runInBand",
			},
			rootPath = "${workspaceFolder}",
			cwd = "${workspaceFolder}",
			console = "integratedTerminal",
			internalConsoleOptions = "neverOpen",
		},
	}
end
