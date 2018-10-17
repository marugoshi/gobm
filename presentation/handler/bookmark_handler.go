package handler

import (
	"context"
	"github.com/marugoshi/gobm/presentation/httputils"
	"github.com/marugoshi/gobm/presentation/service"
)

type BookmarkHandler interface {
	Bookmarks(ctx context.Context, http httputils.Http) error
	Bookmark(ctx context.Context, http httputils.Http) error
}

type bookmarkHandler struct {
	s service.BookmarkService
}

func NewBookmarkHandler(s service.BookmarkService) BookmarkHandler {
	return &bookmarkHandler{s}
}

func (b *bookmarkHandler) Bookmarks(ctx context.Context, http httputils.Http) error {
	return b.s.Bookmarks(ctx, http)
}

func (b *bookmarkHandler) Bookmark(ctx context.Context, http httputils.Http) error {
	return b.s.Bookmark(ctx, http)
}