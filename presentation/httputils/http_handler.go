package httputils

import (
	"net/http"
)

type Params struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

type Handler func(params Params) error
