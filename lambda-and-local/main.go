package main

import (
	handlerLib "github.com/bperezgo/handler-framework"
)

func main() {
	handler := handlerLib.ComposeHandlers(handler1, handler2)
	wrapper(handler.Handle, nil)
}
