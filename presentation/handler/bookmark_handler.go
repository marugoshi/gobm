package handler

import (
	"context"
	"github.com/marugoshi/gobm/domain/service"
	"github.com/marugoshi/gobm/presentation/httputils"
	"os"
	"strconv"
)

type BookmarkHandler interface {
	Bookmarks(ctx context.Context, api httputils.Api) error
	Bookmark(ctx context.Context, api httputils.Api) error
}

type bookmarkHandler struct {
	service.BookmarkService
	prefix string
}

func NewBookmarkHandler(s service.BookmarkService) BookmarkHandler {
	current, _ := os.Getwd()
	return &bookmarkHandler{s, current + "/presentation/view/bookmark"}
}

func (b *bookmarkHandler) Bookmarks(ctx context.Context, api httputils.Api) error {
	data, err := b.BookmarkService.Bookmarks(ctx)
	if err != nil {
		return err
	}
	return api.Html(200, b.prefix+"/index.html", data)
}

func (b *bookmarkHandler) Bookmark(ctx context.Context, api httputils.Api) error {
	id, _ := strconv.Atoi(api.Params[0])
	data, err := b.BookmarkService.Bookmark(ctx, id)
	if err != nil {
		return err
	}
	return api.Html(200, b.prefix+"/show.html", data)
}
