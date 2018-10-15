package httputils

import (
	"fmt"
	"io"
	"net/http"
)

type Params struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

func (p *Params) Text(code int, body string) {
	p.ResponseWriter.Header().Set("Content-Type", ContentTypeTextPlain)
	p.WriteHeader(code)
	io.WriteString(p.ResponseWriter, fmt.Sprintf("%s\n", body))
}

type Func func(params Params) error