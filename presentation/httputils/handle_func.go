package httputils

import (
	"fmt"
	"io"
	"net/http"
)

type Http struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

func (a *Http) Text(code int, body string) {
	a.ResponseWriter.Header().Set("Content-Type", ContentTypeTextPlain)
	a.WriteHeader(code)
	io.WriteString(a.ResponseWriter, fmt.Sprintf("%s\n", body))
}

type Func func(http Http) error