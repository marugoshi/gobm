package httputils

import (
	"net/http"
)

type Params struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

type Func func(params Params) error