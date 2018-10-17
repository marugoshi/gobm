package service

import (
	"context"
	"github.com/marugoshi/gobm/presentation/httputils"
)

type BookmarkService interface {
	Bookmarks(ctx context.Context, http httputils.Http) error
	Bookmark(ctx context.Context, http httputils.Http) error
}

type bookmarkService struct {
}

func NewBookmarkService() BookmarkService {
	return &bookmarkService{}
}

func(b *bookmarkService) Bookmarks(ctx context.Context, http httputils.Http) error {
	return http.RawText(200, "hoge")
}

func(b *bookmarkService) Bookmark(ctx context.Context, http httputils.Http) error {
	return http.RawText(200, http.Params[0])
}