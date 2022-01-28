package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	handlerLib "github.com/bperezgo/handler-framework"
)

type HandlerFunc func(ctx context.Context, req interface{}) (res handlerLib.HandlerResponse)

type WrapperConfig struct {
	Port int
	Body interface{}
	// TODO: Required Parameters, Not defined the functionality
	Method string
	Path   string
}

type localHandlerWrapper struct {
	ctx     context.Context
	handler HandlerFunc
	config  *WrapperConfig
}

func (lh *localHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] Serving Request")
	log.Println("[INFO] Request", r.Body)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(lh.config.Body)
	if err != nil {
		w.WriteHeader(500)
		jsonData := fmt.Sprintf("{\"message\":\"An error ocurred unmarshalling the request\"\"error\":\"%s\"}", err.Error())
		w.Write([]byte(jsonData))
		return
	}
	req := handlerLib.HandlerRequest{
		Body: lh.config.Body,
	}
	res := lh.handler(lh.ctx, req)
	w.WriteHeader(res.StatusCode)
	resByte, _ := json.Marshal(res)
	log.Println("[INFO] Response", string(resByte))
	w.Write([]byte(resByte))
}

func wrapper(handler HandlerFunc, config *WrapperConfig) {
	devEnv, _ := os.LookupEnv("DEV_ENV")
	// Test With DEV_ENV=LOCAL
	if devEnv == "LOCAL" {
		if config == nil {
			log.Fatal("For local use, define config parameter")
		}
		serverUrl := fmt.Sprintf(":%d", config.Port)
		ctx := context.Background()
		localHandler := &localHandlerWrapper{
			ctx:     ctx,
			handler: handler,
			config:  config,
		}
		log.Printf("[INFO] Trying to connect to http://localhost%s%s", serverUrl, config.Path)
		if err := http.ListenAndServe(serverUrl, localHandler); err != nil {
			log.Fatal(err)
		}

		return
	}
	// Test with DEV_ENV=<Anithing or nothing>
	lambda.Start(handler)
}
