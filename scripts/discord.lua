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

t = nil
launched = false
mp.register_event("file-loaded", function()
    if options.active and not launched then
        t = mp.command_native_async({
            name = "subprocess",
            playback_only = false,
            args = { options.binary_path, tostring(pid) }
        })
        launched = true
        msg.info(("(mpv-ipc): %s"):format(socket_path))
    end
end)

mp.register_event("shutdown", function()
    mp.abort_async_command(t)
    os.remove(socket_path) -- finish cleanup
end)
