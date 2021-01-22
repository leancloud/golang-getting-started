package functions

import (
	"github.com/leancloud/go-sdk/leancloud"
)

func init() {
	leancloud.Define("hello", hello)
}

func hello(req *leancloud.FunctionRequest) (interface{}, error) {
	return "Hello World!", nil
}
