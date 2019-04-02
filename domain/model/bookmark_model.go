package model

import (
	"context"
	"github.com/marugoshi/gobm/domain/entity"
)

type BookmarkModel interface {
	All(ctx context.Context, page int, perPage int) ([]*entity.Bookmark, error)
	FindById(ctx context.Context, id int64) (*entity.Bookmark, error)
	Create(ctx context.Context, params *entity.Bookmark) (int64, error)
	Update(ctx context.Context, params *entity.Bookmark) (int64, error)
	Delete(ctx context.Context, id int64) error
}
