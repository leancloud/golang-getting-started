package main

import (
	"github.com/leancloud/go-sdk/leancloud/engine"
)

func registerCloudFunction() {
	engine.Define("hello", func(req *engine.Request) (interface{}, error) {
		return "Hello World!", nil
	})
}
