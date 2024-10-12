<h1 align="center">mpv-discord</h1>

<p align="center">
  <b><small>Discord Rich Presence Integration for MPV Media Player</small></b>
</p>

<p align="center">
  <img alt="1" src="./1.png" width="40%" />
  <img alt="2" src="./2.png" width="40%" />
</p>

<p align="center">
  <sub><b>Left:</b> looping song in playlist (mouse hovering on small icon)</sub>
  <br>
  <sub><b>Right:</b> paused in movie</sub>
</p>

<p align="center">
  <a href="https://github.com/tnychn/mpv-discord/releases"><img alt="github releases" src="https://img.shields.io/github/v/release/tnychn/mpv-discord"></a>
  <a href="https://github.com/tnychn/mpv-discord/releases"><img alt="release date" src="https://img.shields.io/github/release-date/tnychn/mpv-discord"></a>
  <a href="https://github.com/tnychn/mpv-discord/releases"><img alt="downloads" src="https://img.shields.io/github/downloads/tnychn/mpv-discord/total"></a>
  <a href="./LICENSE.txt"><img alt="license" src="https://img.shields.io/github/license/tnychn/mpv-discord"></a>
</p>

## Features

* 🛠 Easy configuration
* 📦 No third-party dependencies
* 🚸 Simple installation (installer scripts included)
* 🏁 Cross-platform (embrace my beloved Golang!)
* ℹ️ Displays song metadata (title, artist, album)
* ⏳ Displays real time player state and timestamps
* 🔕 Toggle activation on the fly by key binding
* 👻 Automatically hide when player is paused
* 🖼 Customizable image assets

## Why?

Currently, there are two alternatives I found on GitHub.

1. [cniw/mpv-discordRPC](https://github.com/cniw/mpv-discordRPC)
2. [noaione/mpv-discordRPC](https://github.com/noaione/mpv-discordRPC)

**Discord RPC**

In order to interact with Discord Rich Presence using RPC, the client needs to connect to Discord's IPC socket.
However, both of the above alternatives do not keep a connection with Discord's IPC socket, which I think is rather unreliable.

See also: [how _mpv-discord_ works](#how-it-works).

**Third-party Dependencies**

Both of the above alternatives require users to install third-party dependencies such as `python-pypresence` or `lua-discordRPC`.
I found it hard to set up the dependencies and I also don't want to mess up my environment.

## Installation

Installer scripts for Windows, Linux and OSX are provided.

1. Download .zip from [the latest release](https://github.com/tnychn/mpv-discord/releases/latest) and extract it.
    * or you can download .zip by clicking on the green download button in GitHub
    * or you can also use `git clone https://github.com/tnychn/mpv-discord.git`
2. Run the installer script of your platform.
    * run `install_darwin.sh` in Terminal for OSX
    * run `install_linux.sh` in Terminal for Linux
    * run `install_windows.bat` by double clicking on it for Windows
3. Before using, you must specify `binary_path` in the config file first.

## Configurations

For OSX and Linux, config file is located in `~/.config/mpv/script-opts`.

For Windows, config file is located in where the `mpv.exe` executable is.

* **key** (default: `D`): key binding to toggle activation on the fly
* **active** (default: `yes`): whether to activate at launch
* **client_id**: specify your own client ID to [customize](#customization) the images shown in Rich Presence
* **binary_path**: full path to mpv-discord's binary file
* **socket_path** (default: `/tmp/mpvsocket`):
  * `use_static_socket_path=yes`: set the full path to the static IPC socket path
  * `use_static_socket_path=no`: set the full path to the directory placing the IPC socket with a *dynamic name*
* **use_static_socket_path** (default: `yes`): whether to use static IPC socket path or *dynamic name* in the path
* **autohide_threshold** (default: `0`): time in seconds before hiding the presence once player is paused (`0` is off)

*dynamic name* is in the format of `mpv-discord-1234` where `1234` will be the PID of the mpv instance.

## Customization

Go to [Discord Developer Portal](https://discord.com/developers/applications),
create an application and upload the following art assets with their corresponding asset keys:
* `mpv`: large image (app logo)
* `play`: small image used when playing
* `pause`: small image used when paused
* `loop`: small image used when playing and looping
* `buffer`: small image used when buffering

Then, set the `client_id` option in the config to the application ID.

You can also find the already provided client ids and their image assets [here](./assets/).

<p align="center">
  <img alt="3" src="./3.png" width="30%" />
  <img alt="4" src="./4.png" width="50%" />
</p>

<p align="center">
  <sub>Using mpv-discord with IINA, a media player based on mpv.</sub>
  <br>
  <sub>Client ID: 834116350884577280</sub>
</p>

## How It Works

This plugin consists of 3 files.

1. [`discord.lua`](./scripts/discord.lua) -- mpv user script
2. [`discord.conf`](./script-opts/discord.conf) -- configuration file
3. [`mpv-discord` binary](./mpv-discord/main.go) -- backend binary of the plugin

When mpv launches, mpv will run all the user scripts including `discord.lua`.
Then `discord.lua` will read the configurations from `discord.conf` and do two things:
(1) create an `input-ipc-server` socket of mpv. (2) start a subprocess of the `mpv-discord` binary.
Then, `mpv-discord` will interact with the `input-ipc-server` to get the player state and properties of mpv.
Finally, `mpv-discord` will update Discord's Rich Presence with the properties through Discord's IPC socket.

## Contributing

If you have any ideas on how to improve this project or if you think there is a lack of features,
feel free to open an issue, or even better, open a pull request. All contributions are welcome!

---

<p align="center">
  <sub><strong>~ crafted with ♥︎ by tnychn ~</strong></sub>
  <br>
  <sub><strong>MIT © 2024 Tony Chan</strong></sub>
</p>
