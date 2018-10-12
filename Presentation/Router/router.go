package Router

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

const (
	ContentTypeTextPlain = "text/plain"
	ContentTypeTextHtml  = "text/html"
)

type Handle func(e *Exchange)
func (f Handle) ServeHTTP(e *Exchange) {
	f(e)
}

type Route struct {
	Pattern string
	Method  string
	Handle
	Re *regexp.Regexp
}

type Router struct {
	Routes       []Route
	ErrorHandler Handle
}

func NewRouter(contentType string) *Router {
	router := &Router{}
	router.ErrorHandler = notFoundErrorHandler(contentType)
	router.Routes = compileRoutes()
	return router
}

func notFoundErrorHandler(contentType string) Handle {
	var errorHandler Handle
	switch contentType {
	case ContentTypeTextPlain:
		errorHandler = func(e *Exchange) {
			e.Text(http.StatusNotFound, "Not Found")
		}
	case ContentTypeTextHtml:
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
			route.Handle(e)
			return
		}
	}
	r.ErrorHandler(e)
}

type Exchange struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

func (e *Exchange) Text(code int, body string) {
	e.ResponseWriter.Header().Set("Content-Type", ContentTypeTextPlain)
	e.WriteHeader(code)
	io.WriteString(e.ResponseWriter, fmt.Sprintf("%s\n", body))
}
