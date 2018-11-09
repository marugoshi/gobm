package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/marugoshi/gobm/domain/data"
	"github.com/marugoshi/gobm/domain/model"
)

var (
	id    int
	url   string
	title string
	memo  string
)

type BookmarkModel struct {
	*sql.DB
}

func NewBookmarkModel(db *sql.DB) model.BookmarkModel {
	return &BookmarkModel{db}
}

func (b *BookmarkModel) All(ctx context.Context, page int, perPage int) (interface{}, error) {
	rows, err := b.DB.Query("SELECT * FROM bookmarks")
	defer rows.Close()
	if err != nil {
		return nil, nil
	}
	records := make([]*data.Bookmark, 0)
	for rows.Next() {
		if err := rows.Scan(&id, &url, &title, &memo); err != nil {
			return nil, nil
		}
		records = append(records, &data.Bookmark{id, url, title, memo})
	}

	data := struct {
		Records []*data.Bookmark
	}{
		records,
	}

	return data, nil
}

func (b *BookmarkModel) FindById(ctx context.Context, id int) (interface{}, error) {
	if err := b.DB.QueryRow("SELECT * FROM bookmarks WHERE id = ?", id).Scan(&id, &url, &title, &memo); err != nil {
		return nil, nil
	}

	record := &data.Bookmark{id, url, title, memo}
	return record, nil
}

func (b *BookmarkModel) Update(ctx context.Context, params *data.Bookmark) (interface{}, error) {
	tx, _ := b.DB.Begin()
	_, err := tx.Exec("UPDATE bookmarks SET url = ?, title = ?, memo = ? WHERE id = ?", params.Url, params.Title, params.Memo, params.Id)
	if err != nil {
		return nil, nil
	}
	tx.Commit()
	return b.FindById(ctx, params.Id)
}