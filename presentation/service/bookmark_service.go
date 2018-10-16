package service

import (
	"context"
	"github.com/marugoshi/gobm/presentation/httputils"
)

type BookmarkService interface {
	Bookmarks(ctx context.Context, http httputils.Http) error
	ShowBookmark(ctx context.Context, http httputils.Http) error
}

type bookmarkService struct {
}

func NewBookmarkService() BookmarkService {
	return &bookmarkService{}
}

func(b *bookmarkService) Bookmarks(ctx context.Context, http httputils.Http) error {
	http.Text(200, "hoge")
	return nil
}

func(b *bookmarkService) ShowBookmark(ctx context.Context, http httputils.Http) error {
	http.Text(200, http.Params[0])
	return nil
}