@echo off
pushd genny
go build -o "../CS-StoryForge.exe" main_win32.go story_forge.go
popd
