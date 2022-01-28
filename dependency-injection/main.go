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
	dummyObject := &ConcreteDummyObject{}
	handler1 := Handler1{
		dummyObject: dummyObject,
	}
	handler := handlerLib.ComposeHandlers(handler1.Handle, handler2)
	res := handler.Handle(ctx, req)
	log.Println("RESPONSE", res)
}
