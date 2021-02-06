msg = require("mp.msg")
opts = require("mp.options")
utils = require("mp.utils")

options = {
    active = true,
    binary_path = "",
    socket_path = "/tmp/mpvsocket",
    use_static_socket_path = true
}
opts.read_options(options, "discord")

if options.binary_path == "" then
    msg.fatal("Missing binary path in config file.")
    os.exit(1)
end

socket_path = options.socket_path
if not options.use_static_socket_path then
    pid = utils.getpid()
    filename = ("mpv-discord-%s"):format(pid)
    if socket_path == "" then
        socket_path = "/tmp/" -- default
    end
    socket_path = utils.join_path(socket_path, filename)
elseif socket_path == "" then
    msg.fatal("Missing socket path in config file.")
    os.exit(1)
end
msg.info(("(mpv-ipc): %s"):format(socket_path))
mp.set_property("input-ipc-server", socket_path)

version = "1.2.1"
msg.info(("mpv-discord v%s by tnychn"):format(version))

t = nil
launched = false
mp.register_event("file-loaded", function()
    if options.active and not launched then
        t = mp.command_native_async({
            name = "subprocess",
            playback_only = false,
            args = { options.binary_path, socket_path }
        }, function()
            msg.info("launched subprocess")
        end)
        launched = true
    end
end)

mp.register_event("shutdown", function()
    mp.abort_async_command(t)
    if not options.use_static_socket_path then
        os.remove(socket_path)
    end
end)
