#!/usr/bin/env bash

cd "$(dirname "$0")" || exit

mpv_dir=$HOME/.config/mpv
scripts_dir=$mpv_dir/scripts
script_opts_dir=$mpv_dir/script-opts

if [ ! -d "$mpv_dir" ]; then
    mkdir -p "$mpv_dir"
fi
if [ ! -d "$scripts_dir" ]; then
    mkdir "$scripts_dir"
fi
if [ ! -d "$script_opts_dir" ]; then
    mkdir "$script_opts_dir"
fi

echo "Copying config [discord.conf]"
cp ./script-opts/discord.conf "$script_opts_dir"

echo "Copying script [discord.lua]"
cp ./scripts/discord.lua "$scripts_dir"

echo "Copying prebuilt linux binary [mpv-discord]"
cp ./bin/linux/mpv-discord "$mpv_dir"/discord

echo
echo "Done! Please manually edit the following option in the config file:"
echo "binary_path=$mpv_dir/discord"
echo
echo "Path to mpv directory: $mpv_dir"
echo "Path to config file: $script_opts_dir/discord.conf"
