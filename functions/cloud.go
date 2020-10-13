package functions

import (
	"github.com/leancloud/go-sdk/leancloud/engine"
)

func init() {
	if err := engine.Define("hello", hello); err != nil {
		panic(err)
	}
}

func hello(req *engine.Request) (interface{}, error) {
	return "Hello World!", nil
}
