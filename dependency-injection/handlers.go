package main

import (
	"log"

	handlerLib "github.com/bperezgo/handler-framework"
)

type Handler1 struct {
	dummyObject IDummyObject
}

func (h1 *Handler1) Handle(req *handlerLib.HandlerRequest, res *handlerLib.HandlerResponse, next handlerLib.NextFunction) {
	log.Println("Info: Prueba de handlers 1")
	message := h1.dummyObject.RunSomething(req)
	log.Println("Returned some mesasge from dummyObject", message)
	res.StatusCode = 200
	res.Message = message
	next()
}

func handler2(req *handlerLib.HandlerRequest, res *handlerLib.HandlerResponse, next handlerLib.NextFunction) {
	log.Println("Info: Prueba de handlers 2")
}
