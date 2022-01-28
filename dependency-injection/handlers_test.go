package main

import (
	"context"
	"log"
	"reflect"
	"testing"

	handlerLib "github.com/bperezgo/handler-framework"
)

func TestHandlers(t *testing.T) {
	testName := "should return some 'Ran another amazing thing from mock of dummyObject' in the message of the response"
	expectedRes := handlerLib.HandlerResponse{
		StatusCode: 200,
		Message:    "Ran another amazing thing from mock of dummyObject",
	}
	t.Run(testName, func(t *testing.T) {
		log.Println("[INFO] TEST NAME:", testName)
		ctx := context.Background()
		req := struct {
			value string
		}{
			value: "value",
		}
		mockDummyObject := &MockDummyObject{}
		handler1 := Handler1{
			dummyObject: mockDummyObject,
		}
		handler := handlerLib.ComposeHandlers(handler1.Handle, handler2)
		gottenRes := handler.Handle(ctx, req)
		if !reflect.DeepEqual(gottenRes, expectedRes) {
			t.Errorf("Gotten response %+v is different to Expected response %+v", gottenRes, expectedRes)
		}
	})
}
