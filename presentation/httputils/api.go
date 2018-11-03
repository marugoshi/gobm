package httputils

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type Api struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

func (a *Api) Html(code int, data interface{}, paths ...string) error {
	body, err := a.parse(data, paths...)
	if err != nil {
		return err
	}
	return a.show(code, body, ContentTypeTextHtml)
}

func (a *Api) parse(data interface{}, paths ...string) (string, error) {
	t, err := template.ParseFiles(paths...)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}

func (a *Api) show(code int, body string, contentType string) error {
	a.ResponseWriter.Header().Set("Content-Type", contentType)
	a.WriteHeader(code)
	_, err := io.WriteString(a.ResponseWriter, fmt.Sprintf("%s\n", body))
	return err
}
