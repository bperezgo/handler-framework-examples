package main

import (
	"log"

	handlerLib "github.com/bperezgo/handler-framework"
)

func handler1(req *handlerLib.HandlerRequest, res *handlerLib.HandlerResponse, next handlerLib.NextFunction) {
	log.Println("Info: Prueba de handlers 1")
	res.StatusCode = 200
	next()
}

func handler2(req *handlerLib.HandlerRequest, res *handlerLib.HandlerResponse, next handlerLib.NextFunction) {
	log.Println("Info: Prueba de handlers 2")
	res.Message = "Message"
}
