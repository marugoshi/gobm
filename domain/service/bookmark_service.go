package service

import (
	"context"
	"github.com/marugoshi/gobm/domain/entity"
	"github.com/marugoshi/gobm/domain/model"
)

type BookmarkService interface {
	Bookmarks(ctx context.Context, page int, perPage int) (interface{}, error)
	Bookmark(ctx context.Context, id int64) (interface{}, error)
	Create(ctx context.Context, bookmark *entity.Bookmark) (interface{}, error)
	Update(ctx context.Context, bookmark *entity.Bookmark) (interface{}, error)
	Delete(ctx context.Context, id int64) error
}

const (
	PER_PAGE = 30
)

type bookmarkService struct {
	model.BookmarkModel
}

func NewBookmarkService(m model.BookmarkModel) BookmarkService {
	return &bookmarkService{m}
}

func (b *bookmarkService) Bookmarks(ctx context.Context, page int, perPage int) (interface{}, error) {
	if page == 0 {
		page = 1
	}

	if perPage == 0 {
		perPage = PER_PAGE
	}

	bookmarks, err := b.BookmarkModel.All(ctx, page, perPage)
	if err != nil {
		return nil, err
	}

	data := struct {
		Records []*entity.Bookmark
	}{
		bookmarks,
	}

	return data, nil
}

func (b *bookmarkService) Bookmark(ctx context.Context, id int64) (interface{}, error) {
	bookmark, err := b.BookmarkModel.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return bookmark, nil
}

func (b *bookmarkService) Create(ctx context.Context, params *entity.Bookmark) (interface{}, error) {
	bookmark, err := b.BookmarkModel.Create(ctx, params)
	if err != nil {
		return nil, err
	}
	return bookmark, nil
}

func (b *bookmarkService) Update(ctx context.Context, params *entity.Bookmark) (interface{}, error) {
	bookmark, err := b.BookmarkModel.Update(ctx, params)
	if err != nil {
		return nil, err
	}
	return bookmark, nil
}

func (b *bookmarkService) Delete(ctx context.Context, id int64) error {
	err := b.BookmarkModel.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
