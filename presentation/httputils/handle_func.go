package httputils

import (
	"net/http"
)

type HandleFuncParams struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

type HandleFunc func(params HandleFuncParams) error
