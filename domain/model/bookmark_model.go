package model

import (
	"context"
)

type BookmarkModel interface {
	All(ctx context.Context, page int, perPage int) (interface{}, error)
	FindById(ctx context.Context, id int) (interface{}, error)
}
