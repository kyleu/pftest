@ECHO OFF
rem Content managed by Project Forge, see [projectforge.md] for details.

rem Starts the app. It doesn't reload on Windows

cd %~dp0\..

@ECHO ON
echo "Windows doesn't allow reloading... sorry"
go.exe build -gcflags "all=-N -l" -o build/debug/projectforge.exe .
build\debug\projectforge.exe -v --addr=0.0.0.0 all projectforge
