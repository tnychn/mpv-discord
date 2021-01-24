msg = require("mp.msg")
opts = require("mp.options")
utils = require("mp.utils")

options = {
    active = true,
    binary_path = ""
}
opts.read_options(options, "discord")

if options.binary_path == "" then
    msg.fatal("Missing binary path in config file.")
    os.exit(1)
end

version = "1.2.1"
msg.info(("mpv-discord v%s by tnychn"):format(version))

pid = utils.getpid()
socket_path = ("/tmp/mpv-discord-%s"):format(pid)
mp.set_property("input-ipc-server", socket_path)

launched = false
mp.register_event("file-loaded", function()
    if options.active and not launched then
--         utils.subprocess_detached({
--             args = { options.binary_path, pid }
--         })
        io.popen(options.binary_path .. " " .. pid)
        launched = true
        msg.info(("(mpv-ipc): %s"):format(socket_path))
    end
end)

mp.register_event("shutdown", function()
    os.remove(socket_path) -- finish cleanup
end)
