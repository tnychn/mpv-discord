msg = require("mp.msg")
opts = require("mp.options")
utils = require("mp.utils")

options = {
    active = true,
    binary_path = "",
    static_ipc = true,
    dynamic_ipc_prefix = ""
}
opts.read_options(options, "discord")

local function get_temp_path()
    local directory_seperator = package.config:match("([^\n]*)\n?")
    local example_temp_file_path = os.tmpname()

    -- remove generated temp file
    pcall(os.remove, example_temp_file_path)

    local seperator_idx = example_temp_file_path:reverse():find(directory_seperator)
    local temp_path_length = #example_temp_file_path - seperator_idx

    return example_temp_file_path:sub(1, temp_path_length)
end

function join_paths(...)
    local arg={...}
    path = ""
    for i,v in ipairs(arg) do
        path = utils.join_path(path, tostring(v))
    end
    return path;
end

tempDir = get_temp_path()
os.execute("mkdir " .. join_paths(tempDir, options.dynamic_ipc_prefix) .. " 2>/dev/null")

if options.binary_path == "" then
    msg.fatal("Missing binary path in config file.")
    os.exit(1)
end

version = "1.2.1"
msg.info(("mpv-discord v%s by tnychn"):format(version))

pid = utils.getpid()
sep = package.config:sub(1,1)
socket_path = ""
static_ipc_str = 'false'

if sep == '/' then
    if options.static_ipc then
        static_ipc_str = "true"
        socket_path = ("/tmp/mpvsocket")
    else
        socket_path = join_paths(tempDir, options.dynamic_ipc_prefix, pid)
    end
    
else
    if options.static_ipc then
        static_ipc_str = "true"
        socket_path = ("mpvpipe")
    else
        socket_path = join_paths(tempDir, options.dynamic_ipc_prefix, pid)
    end
    
end
mp.set_property("input-ipc-server", socket_path)


if options.static_ipc then
    static_ipc_str = "true"
end

launched = false
mp.register_event("file-loaded", function()
    if options.active and not launched then
--         utils.subprocess_detached({
--             args = { options.binary_path, pid }
--         })
        io.popen(options.binary_path .. " " .. socket_path)
        launched = true
        msg.info(("(mpv-ipc): %s"):format(socket_path))
    end
end)

mp.register_event("shutdown", function()
    os.remove(socket_path) -- finish cleanup
end)
