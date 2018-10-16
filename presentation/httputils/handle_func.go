package httputils

import (
	"context"
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
	a.show(code, body, ContentTypeTextPlain)
}

func (a *Http) Html(code int, body string) {
	a.show(code, body, ContentTypeTextHtml)
}

func (a *Http) show(code int, body string, contentType string) {
	a.ResponseWriter.Header().Set("Content-Type", contentType)
	a.WriteHeader(code)
	io.WriteString(a.ResponseWriter, fmt.Sprintf("%s\n", body))
}

type Func func(ctx context.Context, http Http) error
