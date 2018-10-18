package handler

import (
	"github.com/marugoshi/gobm/domain/service"
)

type Handlers interface {
	BookmarkHandler
}

type handlers struct {
	BookmarkHandler
}

func NewHandlers() Handlers {
	return &handlers{NewBookmarkHandler(service.NewBookmarkService())}
}
