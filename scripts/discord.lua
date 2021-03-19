msg = require("mp.msg")
opts = require("mp.options")
utils = require("mp.utils")

options = {
    key = "D",
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

version = "1.4.1"
msg.info(("mpv-discord v%s by tnychn"):format(version))

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

cmd = nil

function start()
    if options.active and cmd == nil then
        cmd = mp.command_native_async({
            name = "subprocess",
            playback_only = false,
            args = { options.binary_path, socket_path }
        }, function()
            msg.info("launched subprocess")
        end)
        mp.osd_message("Discord Rich Presence: Started")
    end
end

function stop()
    mp.abort_async_command(cmd)
    cmd = nil
    msg.info("aborted subprocess")
    mp.osd_message("Discord Rich Presence: Stopped")
end

mp.register_event("file-loaded", start)

mp.add_key_binding(options.key, "toggle-discord", function()
    if cmd ~= nil then
        stop()
    else
        start()
    end
end)

mp.register_event("shutdown", function()
    if cmd ~= nil then
        stop()
    end
    if not options.use_static_socket_path then
        os.remove(socket_path)
    end
end)
