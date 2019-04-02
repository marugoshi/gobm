package model

import (
	. "context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/marugoshi/gobm/domain/entity"
	"github.com/marugoshi/gobm/domain/model"
	. "github.com/marugoshi/gobm/infrastructure/storage/mysql"
	"github.com/pkg/errors"
)

type BookmarkModel struct {
	db *sql.DB
}

func NewBookmarkModel(db *sql.DB) model.BookmarkModel {
	return &BookmarkModel{db: db}
}

func (b *BookmarkModel) All(ctx Context, page int, perPage int) ([]*entity.Bookmark, error) {
	q := `
SELECT
	id, directory_id, url, title, created_at, updated_at
FROM
	bookmarks
LIMIT
	?
OFFSET
	?
`

	rows, err := b.db.Query(q, perPage, (page-1)*perPage)
	defer rows.Close()
	if err != nil {
		return nil, errors.Wrap(err, "can not execute query.")
	}

	records := make([]*entity.Bookmark, 0)
	for rows.Next() {
		bookmark := &entity.Bookmark{}
		if err := rows.Scan(
			&bookmark.Id,
			&bookmark.DirectoryId,
			&bookmark.Url,
			&bookmark.Title,
			&bookmark.CreatedAt,
			&bookmark.UpdatedAt); err != nil {
			return nil, errors.Wrap(err, "can not scan.")
		}
		records = append(records, bookmark)
	}

	return records, nil
}

func (b *BookmarkModel) FindById(ctx Context, id int64) (*entity.Bookmark, error) {
	q := `
SELECT 
	id, directory_id, url, title, created_at, updated_at
FROM
	bookmarks
WHERE
	id = ?
`

	bookmark := &entity.Bookmark{}
	if err := b.db.QueryRow(q, id).Scan(
		&bookmark.Id,
		&bookmark.DirectoryId,
		&bookmark.Url,
		&bookmark.Title,
		&bookmark.CreatedAt,
		&bookmark.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, errors.Wrap(err, "can not execute query.")
		}
	}

	return bookmark, nil
}

func (b *BookmarkModel) Create(ctx Context, params *entity.Bookmark) (int64, error) {
	q := `
INSERT INTO
	bookmarks
(
	url,
	directory_id,
	title,
	created_at,
	updated_at
) VALUES(
	?, ?, ?, NOW(), NOW()
)
`

	result, err := b.db.Exec(q,
		params.Url,
		params.DirectoryId,
		params.Title)
	if err != nil {
		return 0, errors.Wrap(err, "can not execute query.")
	}

	return result.LastInsertId()

}

func (b *BookmarkModel) Update(ctx Context, params *entity.Bookmark) (int64, error) {
	q := `
UPDATE
	bookmarks
SET
	directory_id = ?,
	url = ?,
	title = ?,
	updated_at = NOW()
WHERE
	id = ?
`

	if err := Transaction(b.db, func(tx *sql.Tx) error {
		_, err := tx.Exec(q,
			params.DirectoryId,
			params.Url,
			params.Title,
			params.Id)
		if err != nil {
			return err
		}
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "can not execute query.")
	}

	return params.Id, nil
}

func (b *BookmarkModel) Delete(ctx Context, id int64) error {
	q := `
DELETE FROM
	bookmarks
WHERE
	id = ?
`

	if err := Transaction(b.db, func(tx *sql.Tx) error {
		_, err := tx.Exec(q, id)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "can not execute query.")
	}

	return nil
}
