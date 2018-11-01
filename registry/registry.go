package registry

import (
	"github.com/marugoshi/gobm/domain/service"
	"github.com/marugoshi/gobm/presentation/handler"
)

type Registry struct {
	service.BookmarkService
	handler.BookmarkHandler
}

func NewRegistry() Registry {
	bookmarkService := service.NewBookmarkService()
	bookmarkHandler := handler.NewBookmarkHandler(bookmarkService)
	return Registry{bookmarkService, bookmarkHandler}
}
