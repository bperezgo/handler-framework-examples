package main

import (
	handlerLib "github.com/bperezgo/handler-framework"
)

type RequestBody struct {
	Value string `json:"value"`
}

func main() {
	handler := handlerLib.ComposeHandlers(handler1, handler2)
	config := &WrapperConfig{
		Port: 3000,
		Body: &RequestBody{},
		// TODO: This path doesnot work yet
		Path: "/",
		// TODO: This method doesnot work yet
		Method: "POST",
	}
	wrapper(handler.Handle, config)
}
