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
}

type localHandlerWrapper struct {
	ctx     context.Context
	handler HandlerFunc
}

func (lh *localHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := handlerLib.HandlerRequest{
		Body: r.Body,
	}
	res := lh.handler(lh.ctx, req)
	w.WriteHeader(res.StatusCode)
	resByte, _ := json.Marshal(res)
	w.Write([]byte(resByte))
}

func wrapper(handler HandlerFunc, config *WrapperConfig) {
	devEnv, _ := os.LookupEnv("DEV_ENV")
	// Test With DEV_ENV=LOCAL
	if devEnv == "LOCAL" {
		serverUrl := fmt.Sprintf(":%d", config.Port)
		ctx := context.Background()
		localHandler := &localHandlerWrapper{
			ctx:     ctx,
			handler: handler,
		}
		if err := http.ListenAndServe(serverUrl, localHandler); err != nil {
			log.Fatal(err)
		}

		return
	}
	// Test with DEV_ENV=<Anithing or nothing>
	lambda.Start(handler)
}
