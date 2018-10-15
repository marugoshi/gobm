package router

import (
	"github.com/marugoshi/gobm/presentation/httputils"
	"regexp"
)

type Routes interface {
	Data() []Route
}

type routes struct {
	data []Route
}

func (r *routes) Data() []Route {
	return r.data
}

func NewRoutes() Routes {
	routes := &routes{}
	return routes
}

type Route struct {
	Pattern *regexp.Regexp
	Method  string
	httputils.Func
}