package main

import (
	"fmt"
	"syscall/js"
)

var htmlString = "<h4>Button clicked</h4>"

func GetHtml() js.Func {
	return js.FuncOf(func(this js.Value, arg []js.Value) interface{} {
		return htmlString
	})
}

func main() {
	ch := make(chan struct{}, 0)

	fmt.Print("hello, world\n")

	js.Global().Set("getHtml", GetHtml())

	<-ch
}
