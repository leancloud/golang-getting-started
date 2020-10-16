package functions

import (
	"github.com/leancloud/go-sdk/leancloud/engine"
)

func init() {
	engine.Define("hello", hello)
}

func hello(req *engine.Request) (interface{}, error) {
	return "Hello World!", nil
}
