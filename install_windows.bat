@echo off
title "mpv-discord installer script"

set src_dir=%~dp0
cd %src_dir%

:set_mpv_dir
echo "Enter full path to the mpv folder (where mpv.exe is located)"
set /p mpv_dir="mpv folder: "

if exist "%mpv_dir%" (
	if exist "%mpv_dir%\mpv.exe" (
		goto install
	) else (
		echo "Failed to locate 'mpv.exe' in '%mpv_dir%'"
		exit /b 1
	)
) else (
	echo "Path does not exist"
    exit /b 1
)

:install
set scripts_dir="%mpv_dir%\scripts"
set script_opts_dir="%mpv_dir%\script-opts"
if not exist "%scripts_dir%" mkdir "%scripts_dir%"
if not exist "%script_opts_dir%" mkdir "%script_opts_dir%"

echo "Copying config [discord.conf]"
copy "%src_dir%\script-opts\discord.conf" "%script_opts_dir%" > nul

echo "Copying scripts [discord.lua]"
copy "%src_dir%\scripts\discord.lua" "%scripts_dir%" > nul

echo "Copying prebuilt windows binary [mpv-discord.exe]"
copy "%src_dir%\bin\windows\mpv-discord.exe" "%mpv_dir%\discord.exe" > nul

echo:
echo "Done! Please manually edit the following option in the config file:"
echo "binary_path=%mpv_dir%\discord.exe"
echo:
echo "Path to mpv directory: %mpv_dir%"
echo "Path to config file: %script_opts_dir%/discord.conf"
