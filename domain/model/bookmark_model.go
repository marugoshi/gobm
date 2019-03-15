package model

import (
	"context"
	"github.com/marugoshi/gobm/domain/entity"
)

type BookmarkModel interface {
	All(ctx context.Context, page int, perPage int) (interface{}, error)
	FindById(ctx context.Context, id int64) (interface{}, error)
	Create(ctx context.Context, params *entity.Bookmark) (interface{}, error)
	Update(ctx context.Context, params *entity.Bookmark) (interface{}, error)
	Delete(ctx context.Context, id int64) error
}
