package router

import (
	"github.com/marugoshi/gobm/presentation/httputils"
	"net/http"
	"regexp"
)

type HandleFunc func(res http.ResponseWriter, req *http.Request, params []string) error

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
		errorHandler = func(res http.ResponseWriter, req *http.Request, params []string) error {
			return nil
			// e.Text(http.StatusNotFound, "Not Found")
		}
	case httputils.ContentTypeTextHtml:
		errorHandler = func(res http.ResponseWriter, req *http.Request, params []string) error {
			return nil
		}
	default:
		errorHandler = func(res http.ResponseWriter, req *http.Request, params []string) error {
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
	params := make([]string, 0)
	for _, route := range r.Routes {
		if matches := route.Re.FindStringSubmatch(req.URL.Path); len(matches) > 0 && route.Method == req.Method {
			if len(matches) > 1 {
				params = matches[1:]
			}
			route.HandleFunc(res, req, params)
			return
		}
	}
	r.ErrorHandler(res, req, params)
}

/*
// ハンドラーが持てばよいのでは？
func (e *Exchange) Text(code int, body string) {
	e.ResponseWriter.Header().Set("Content-Type", httputils.ContentTypeTextPlain)
	e.WriteHeader(code)
	io.WriteString(e.ResponseWriter, fmt.Sprintf("%s\n", body))
}
*/
