package router

import (
	"github.com/marugoshi/gobm/presentation/httputils"
	"net/http"
)

type Router struct {
	Routes
	httputils.Func
}

func NewRouter(contentType string) *Router {
	return &Router{NewRoutes(), notFoundError(contentType)}
}

func notFoundError(contentType string) httputils.Func {
	var errorHandler httputils.Func
	switch contentType {
	case httputils.ContentTypeTextPlain:
		errorHandler = func(params httputils.Params) error {
			return nil
			// e.Text(http.StatusNotFound, "Not Found")
		}
	case httputils.ContentTypeTextHtml:
		errorHandler = func(params httputils.Params) error {
			return nil
		}
	default:
		errorHandler = func(params httputils.Params) error {
			return nil
		}
	}
	return errorHandler
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	params := httputils.Params{ResponseWriter: res, Request: req}
	for _, route := range r.Routes.Data() {
		if matches := route.Pattern.FindStringSubmatch(req.URL.Path); len(matches) > 0 && route.Method == req.Method {
			if len(matches) > 1 {
				params.Params = matches[1:]
			}
			route.Func(params)
			return
		}
	}
	r.Func(params)
}

/*
// ハンドラーが持てばよいのでは？
func (e *Exchange) Text(code int, body string) {
	e.ResponseWriter.Header().Set("Content-Type", httputils.ContentTypeTextPlain)
	e.WriteHeader(code)
	io.WriteString(e.ResponseWriter, fmt.Sprintf("%s\n", body))
}
*/
