return {
  "rcarriga/nvim-notify",
  lazy = false,
  priority = 990,
  config = function()
    -- Sets up the notify settings for the notification modal that will be used for showing info, warnings, and errors in the config
    local notify = require("notify")
    notify.setup({
      ---@usage Animation style one of { "fade", "slide", "fade_in_slide_out", "static" }
      stages = "slide",
      ---@usage Function called when a new window is opened, use for changing win settings/config
      on_open = nil,
      ---@usage Function called when a window is closed
      on_close = nil,
      ---@usage timeout for notifications in ms, default 5000
      timeout = 10000,
      -- Render function for notifications. See notify-render()
      render = "compact",
      ---@usage highlight behind the window for stages that change opacity
      background_colour = "Normal",
      ---@usage minimum width for notification windows
      minimum_width = 50,
      ---@usage Icons for the different levels
      icons = {
        ERROR = "",
        WARN = "",
        INFO = "",
        DEBUG = "",
        TRACE = "✎",
      },
    })

    vim.notify = notify
  end
}