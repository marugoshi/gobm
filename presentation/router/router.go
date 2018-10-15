package router

import (
	"github.com/marugoshi/gobm/presentation/handler"
	"github.com/marugoshi/gobm/presentation/httputils"
	"net/http"
	"regexp"
)

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
	var routesData = []Route{
		Route{`^/bookmarks$`, http.MethodGet, func(params httputils.Params) error {
			return r.Handlers.Bookmarks(params)
		}},
	}

	params := httputils.Params{ResponseWriter: res, Request: req}
	for _, route := range routesData {
		re := regexp.MustCompile(route.Pattern)
		if matches := re.FindStringSubmatch(req.URL.Path); len(matches) > 0 && route.Method == req.Method {
			if len(matches) > 1 {
				params.Params = matches[1:]
			}
			route.Func(params)
			return
		}
	}
	r.Func(params)
}