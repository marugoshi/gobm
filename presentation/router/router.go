package router

import (
	"fmt"
	"github.com/marugoshi/gobm/presentation/httputils"
	"io"
	"net/http"
	"regexp"
)

type HandleFunc func(e *Exchange)

type Route struct {
	Pattern string
	Method  string
	HandleFunc
	Re *regexp.Regexp
}

type Router struct {
	Routes       []Route
	ErrorHandler HandleFunc
}

func NewRouter(contentType string) *Router {
	router := &Router{}
	router.ErrorHandler = notFoundErrorHandler(contentType)
	router.Routes = compileRoutes()
	return router
}

func notFoundErrorHandler(contentType string) HandleFunc {
	var errorHandler HandleFunc
	switch contentType {
	case httputils.ContentTypeTextPlain:
		errorHandler = func(e *Exchange) {
			e.Text(http.StatusNotFound, "Not Found")
		}
	case httputils.ContentTypeTextHtml:
		errorHandler = func(e *Exchange) {
		}
	default:
		errorHandler = func(e *Exchange) {
		}
	}
	return errorHandler
}

func compileRoutes() []Route {
	results := make([]Route, 0)
	for _, route := range Routes {
		route.Re = regexp.MustCompile(route.Pattern)
		results = append(results, route)
	}
	return results
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	e := &Exchange{ResponseWriter: res, Request: req}
	for _, route := range r.Routes {
		if matches := route.Re.FindStringSubmatch(e.URL.Path); len(matches) > 0 && route.Method == req.Method {
			if len(matches) > 1 {
				e.Params = matches[1:]
			}
			route.HandleFunc(e)
			return
		}
	}
	r.ErrorHandler(e)
}

// 消したい
type Exchange struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

// ハンドラーが持てばよいのでは？
func (e *Exchange) Text(code int, body string) {
	e.ResponseWriter.Header().Set("Content-Type", httputils.ContentTypeTextPlain)
	e.WriteHeader(code)
	io.WriteString(e.ResponseWriter, fmt.Sprintf("%s\n", body))
}
