package entity

import (
	"github.com/marugoshi/gobm/infrastructure/storage/mysql"
)

type Bookmark struct {
	Id          int64
	DirectoryId mysql.NullInt64
	Url         string
	Title       string
	CreatedAt   mysql.NullDateTime
	UpdatedAt   mysql.NullDateTime
}
