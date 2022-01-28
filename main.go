package main

import (
	"context"
	"log"

	handlerLib "github.com/bperezgo/handler-framework"
)

func main() {
	ctx := context.Background()
	req := struct {
		value string
	}{
		value: "value",
	}
	handler := handlerLib.ComposeHandlers(handler1, handler2)
	res := handler.Handle(ctx, req)
	log.Println("RESPONSE", res)
}
