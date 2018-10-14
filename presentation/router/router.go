package router

import (
	"github.com/marugoshi/gobm/presentation/httputils"
	"net/http"
	"regexp"
)

type Route struct {
	Pattern string
	Method  string
	httputils.HandleFunc
	Re *regexp.Regexp
}

type Router struct {
	Routes       []Route
	httputils.HandleFunc
}

func NewRouter(contentType string) *Router {
	router := &Router{}
	router.HandleFunc = notFoundErrorHandler(contentType)
	router.Routes = compileRoutes()
	return router
}

func notFoundErrorHandler(contentType string) httputils.HandleFunc {
	var errorHandler httputils.HandleFunc
	switch contentType {
	case httputils.ContentTypeTextPlain:
		errorHandler = func(params httputils.HandleFuncParams) error {
			return nil
			// e.Text(http.StatusNotFound, "Not Found")
		}
	case httputils.ContentTypeTextHtml:
		errorHandler = func(params httputils.HandleFuncParams) error {
			return nil
		}
	default:
		errorHandler = func(params httputils.HandleFuncParams) error {
			return nil
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
	handleFuncParams := httputils.HandleFuncParams{ResponseWriter: res, Request: req}
	for _, route := range r.Routes {
		if matches := route.Re.FindStringSubmatch(req.URL.Path); len(matches) > 0 && route.Method == req.Method {
			if len(matches) > 1 {
				handleFuncParams.Params = matches[1:]
			}
			route.HandleFunc(handleFuncParams)
			return
		}
	}
	r.HandleFunc(handleFuncParams)
}

/*
// ハンドラーが持てばよいのでは？
func (e *Exchange) Text(code int, body string) {
	e.ResponseWriter.Header().Set("Content-Type", httputils.ContentTypeTextPlain)
	e.WriteHeader(code)
	io.WriteString(e.ResponseWriter, fmt.Sprintf("%s\n", body))
}
*/
