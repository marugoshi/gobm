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
		errorHandler = func(api httputils.Api) error {
			return nil
			// e.Text(http.StatusNotFound, "Not Found")
		}
	case httputils.ContentTypeTextHtml:
		errorHandler = func(api httputils.Api) error {
			return nil
		}
	default:
		errorHandler = func(api httputils.Api) error {
			return nil
		}
	}
	return errorHandler
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var routesData = []Route{
		Route{`^/bookmarks$`, http.MethodGet, func(api httputils.Api) error {
			return r.Handlers.Bookmarks(api)
		}},
	}

	api := httputils.Api{ResponseWriter: res, Request: req}
	for _, route := range routesData {
		re := regexp.MustCompile(route.Pattern)
		if matches := re.FindStringSubmatch(req.URL.Path); len(matches) > 0 && route.Method == req.Method {
			if len(matches) > 1 {
				api.Params = matches[1:]
			}
			route.Func(api)
			return
		}
	}
	r.Func(api)
}