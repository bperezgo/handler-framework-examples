package main

import (
	"log"

	handlerLib "github.com/bperezgo/handler-framework"
)

func handler1(req *handlerLib.HandlerRequest, res *handlerLib.HandlerResponse, next handlerLib.NextFunction) {
	log.Println("[INFO] This will run in a lambda or in local")
	next()
}

func handler2(req *handlerLib.HandlerRequest, res *handlerLib.HandlerResponse, next handlerLib.NextFunction) {
	log.Println("[INFO] The last handler")
	res.StatusCode = 200
	res.Message = "Last handler"
}
