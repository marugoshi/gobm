package handler

import (
	"context"
	"github.com/marugoshi/gobm/presentation/httputils"
	"github.com/marugoshi/gobm/domain/service"
	"os"
	"strconv"
)

type BookmarkHandler interface {
	Bookmarks(ctx context.Context, http httputils.Http) error
	Bookmark(ctx context.Context, http httputils.Http) error
}

type bookmarkHandler struct {
	s service.BookmarkService
	prefix string
}

func NewBookmarkHandler(s service.BookmarkService) BookmarkHandler {
	current, _ := os.Getwd()
	return &bookmarkHandler{s, current + "/presentation/view/bookmark"}
}

func (b *bookmarkHandler) Bookmarks(ctx context.Context, http httputils.Http) error {
	data, err := b.s.Bookmarks(ctx)
	if err != nil {
		return err
	}
	return http.Html(200, "index", b.prefix + "/index.html", data)
}

func (b *bookmarkHandler) Bookmark(ctx context.Context, http httputils.Http) error {
	id, _ := strconv.Atoi(http.Params[0])
	data, err := b.s.Bookmark(ctx, id)
	if err != nil {
		return err
	}
	return http.Html(200, "index", b.prefix + "/show.html", data)
}