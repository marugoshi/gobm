package entity

import (
	"database/sql"
)

type Bookmark struct {
	Id          int64
	DirectoryId sql.NullInt64
	Url         string
	Title       string
}
