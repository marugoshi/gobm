package router

import (
	"github.com/marugoshi/gobm/presentation/handler"
	"github.com/marugoshi/gobm/presentation/httputils"
	"net/http"
	"regexp"
)

func (r *Router) routesData() []Route {
	return []Route{
		Route{`^/bookmarks$`, http.MethodGet, r.Handlers.Bookmarks },
	}
}

type Route struct {
	Pattern string
	Method  string
	httputils.Func
}

type Router struct {
	handler.Handlers
	httputils.Func
}

func NewRouter(contentType string) *Router {
	return &Router{handler.NewHandlers(), notFoundError(contentType)}
}

func notFoundError(contentType string) httputils.Func {
	var errorHandler httputils.Func
	switch contentType {
	case httputils.ContentTypeTextPlain:
		errorHandler = func(http httputils.Http) error {
			return nil
			// e.Text(http.StatusNotFound, "Not Found")
		}
	case httputils.ContentTypeTextHtml:
		errorHandler = func(http httputils.Http) error {
			return nil
		}
	default:
		errorHandler = func(http httputils.Http) error {
			return nil
		}
	}
	return errorHandler
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	http := httputils.Http{ResponseWriter: res, Request: req}
	for _, route := range r.routesData() {
		re := regexp.MustCompile(route.Pattern)
		if matches := re.FindStringSubmatch(req.URL.Path); len(matches) > 0 && route.Method == req.Method {
			if len(matches) > 1 {
				http.Params = matches[1:]
			}
			route.Func(http)
			return
		}
	}
	r.Func(http)
}