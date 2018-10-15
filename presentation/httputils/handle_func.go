package httputils

import (
	"fmt"
	"io"
	"net/http"
)

type Api struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

func (a *Api) Text(code int, body string) {
	a.ResponseWriter.Header().Set("Content-Type", ContentTypeTextPlain)
	a.WriteHeader(code)
	io.WriteString(a.ResponseWriter, fmt.Sprintf("%s\n", body))
}

type Func func(api Api) error