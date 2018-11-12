package router

import (
	"context"
	"github.com/marugoshi/gobm/presentation/httputils"
	"github.com/marugoshi/gobm/registry"
	"log"
	"net/http"
	"regexp"
)

func (r *Router) routesData() []Route {
	return []Route{
		Route{`^/bookmarks$`, http.MethodGet, r.Registry.BookmarkIndex},
		Route{`^/bookmarks/new$`, http.MethodGet, r.Registry.BookmarkNew},
		Route{`^/bookmarks$`, http.MethodPost, r.Registry.BookmarkCreate},
		Route{`^/bookmarks/(\d*)/edit$`, http.MethodGet, r.Registry.BookmarkEdit},
		Route{`^/bookmarks/(\d*)$`, http.MethodPatch, r.Registry.BookmarkUpdate},
		Route{`^/bookmarks/(\d*)$`, http.MethodDelete, r.Registry.BookmarkDelete},
	}
}

type Route struct {
	Pattern string
	Method  string
	httputils.Func
}

type Router struct {
	registry.Registry
	NotFound httputils.Func
}

func NewRouter(registry registry.Registry, contentType string) *Router {
	return &Router{registry, notFoundError(contentType)}
}

func notFoundError(contentType string) httputils.Func {
	var errorHandler httputils.Func
	switch contentType {
	case httputils.ContentTypeTextHtml:
		errorHandler = notFoundErrorTextHtml
	default:
		errorHandler = notFoundErrorTextHtml
	}
	return errorHandler
}

func notFoundErrorTextHtml(ctx context.Context, api httputils.Api) error {
	return nil
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	api := httputils.Api{ResponseWriter: res, Request: req}

	if !r.serveHTTP(ctx, api, r.routesData()) {
		if !r.serveStaticHTTP(res, req) {
			r.NotFound(ctx, api)
		}
	}
}

func (r *Router) serveHTTP(ctx context.Context, api httputils.Api, routes []Route) bool {
	method := r.normalizeMethod(api.Request)
	for _, route := range routes {
		re := regexp.MustCompile(route.Pattern)
		if matches := re.FindStringSubmatch(api.Request.URL.Path); len(matches) > 0 && route.Method == method {
			if len(matches) > 1 {
				api.Params = matches[1:]
			}
			err := route.Func(ctx, api)
			if err != nil {
				log.Fatal(err)
				return true
			} else {
				return true
			}
		}
	}
	return false
}

func (r *Router) normalizeMethod(req *http.Request) string {
	method := req.FormValue("_method")
	if method == "" {
		method = req.Method
	}
	return method
}

func (r *Router) serveStaticHTTP(res http.ResponseWriter, req *http.Request) bool {
	staticRe := regexp.MustCompile(`^/static`)
	if staticMatch := staticRe.MatchString(req.URL.Path); staticMatch {
		http.ServeFile(res, req, req.URL.Path[1:])
		return true
	}
	return false
}
