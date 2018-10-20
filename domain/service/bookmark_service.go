package service

import (
	"context"
)

type BookmarkService interface {
	Bookmarks(ctx context.Context) (interface{}, error)
	Bookmark(ctx context.Context, id int) (interface{}, error)
}

type bookmarkService struct{}

func NewBookmarkService() BookmarkService {
	return &bookmarkService{}
}

func (b *bookmarkService) Bookmarks(ctx context.Context) (interface{}, error) {
	data := struct {
		Key string
	}{
		Key: "hoge",
	}
	return data, nil
}

func (b *bookmarkService) Bookmark(ctx context.Context, id int) (interface{}, error) {
	data := struct {
		Key string
	}{
		Key: "hoge",
	}
	return data, nil
}
