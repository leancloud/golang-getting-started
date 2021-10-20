package functions

import (
	"github.com/leancloud/go-sdk/leancloud"
)

func init() {
	leancloud.Engine.Define("hello", hello)
}

func hello(req *leancloud.FunctionRequest) (interface{}, error) {
	return map[string]string{
		"hello": "world",
	}, nil
}
