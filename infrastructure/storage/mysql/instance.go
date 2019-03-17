package mysql

import (
	"database/sql"
	"fmt"
	"os"
)

func NewInstance() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&loc=Asia%%2FTokyo", user, pass, host, name))
}
