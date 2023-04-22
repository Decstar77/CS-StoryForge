@echo off
set GOOS=js
set GOARCH=wasm

pushd genny
go build -o "../site/cs-story-forge.wasm" main_wasm.go story_forge.go
popd

