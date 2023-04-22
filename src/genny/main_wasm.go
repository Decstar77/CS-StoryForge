package main

import (
	"bytes"
	"fmt"
	"io"
	"syscall/js"
)

func Go_GenerateProompt() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return "Error: Missing argument for 'fileData'"
		}

		fileData := args[0]
		length := fileData.Get("byteLength").Int()
		data := make([]byte, length)
		js.CopyBytesToGo(data, fileData)

		reader := io.Reader(bytes.NewReader(data))

		proompt := GenerateProompt(reader)

		return proompt.text
	})
}

func DoWebApp() {
	ch := make(chan struct{}, 0)
	fmt.Printf("Hello Web Assembly from Go!\n")

	js.Global().Set("Go_GenerateProompt", Go_GenerateProompt())
	<-ch
}

func main() {
	DoWebApp()
}
