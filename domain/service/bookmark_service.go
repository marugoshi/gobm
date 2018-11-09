package service

import (
	"context"
	"github.com/marugoshi/gobm/domain/data"
	"github.com/marugoshi/gobm/domain/model"
)

type BookmarkService interface {
	Bookmarks(ctx context.Context) (interface{}, error)
	Bookmark(ctx context.Context, id int64) (interface{}, error)
	Create(ctx context.Context, bookmark *data.Bookmark) (interface{}, error)
	Update(ctx context.Context, bookmark *data.Bookmark) (interface{}, error)
}

type bookmarkService struct {
	model.BookmarkModel
}

func NewBookmarkService(m model.BookmarkModel) BookmarkService {
	return &bookmarkService{m}
}

func (b *bookmarkService) Bookmarks(ctx context.Context) (interface{}, error) {
	bookmarks, err := b.BookmarkModel.All(ctx, 1, 100)
	if err != nil {
		return nil, nil
	}
	return bookmarks, nil
}

func (b *bookmarkService) Bookmark(ctx context.Context, id int64) (interface{}, error) {
	bookmark, err := b.BookmarkModel.FindById(ctx, id)
	if err != nil {
		return nil, nil
	}
	return bookmark, nil
}

func (b *bookmarkService) Create(ctx context.Context, params *data.Bookmark) (interface{}, error) {
	bookmark, err := b.BookmarkModel.Create(ctx, params)
	if err != nil {
		return nil, nil
	}
	return bookmark, nil
}

func (b *bookmarkService) Update(ctx context.Context, params *data.Bookmark) (interface{}, error) {
	bookmark, err := b.BookmarkModel.Update(ctx, params)
	if err != nil {
		return nil, nil
	}
	return bookmark, nil
}
