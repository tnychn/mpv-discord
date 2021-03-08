@echo off

set src_dir=%~dp0
cd %src_dir%

set /p mpv_dir=Path: 

if exist %mpv_dir% (
    if not exist %mpv_dir%\mpv.exe (
        echo Failed to locate 'mpv.exe' in '%mpv_dir%'.
        exit /b 1
    )
) else (
    echo Path '%mpv_dir%' does not exist.
    exit /b 1
)

set config_dir=%mpv_dir%
if exist %mpv_dir%\portable_config set config_dir=%mpv_dir%\portable_config

set scripts_dir=%config_dir%\scripts
set script_opts_dir=%config_dir%\script-opts
if not exist %scripts_dir% mkdir %scripts_dir%
if not exist %script_opts_dir% mkdir %script_opts_dir%

echo Copying [discord.conf]
copy %src_dir%\script-opts\discord.conf %script_opts_dir% >nul

echo Copying [discord.lua]
copy %src_dir%\scripts\discord.lua %scripts_dir% >nul

echo Copying [mpv-discord.exe]
copy %src_dir%\bin\windows\mpv-discord.exe %config_dir%\discord.exe >nul

echo:
echo Path to mpv directory: %mpv_dir%
echo Path to config file: %script_opts_dir%/discord.conf
echo:
echo Please manually edit the following option in the config file:
echo:
echo   binary_path=%config_dir%\discord.exe
echo:
