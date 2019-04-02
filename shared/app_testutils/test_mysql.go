package app_testutils

import (
	"database/sql"
	"fmt"
	"github.com/marugoshi/gobm/infrastructure/storage/mysql"
	"os"
	"os/exec"
	"testing"
)

func Setup() {
	os.Setenv("DB_NAME", "gobm_t")
}

func Teardown() {
	os.Setenv("DB_NAME", "gobm_d")
}

type TestMysql interface {
	GetInstance() *sql.DB
	Fixtures(queries []string)
	Truncates(tables []string)
}

type testMysql struct {
	t  *testing.T
	db *sql.DB
}

func NewTestMySql(t *testing.T) TestMysql {
	db, err := mysql.NewInstance()
	if err != nil {
		t.Fatalf("can not create db instance: %v", err)
	}
	return &testMysql{t: t, db: db}
}

func (m testMysql) GetInstance() *sql.DB {
	return m.db
}

func (m testMysql) Fixtures(queries []string) {
	password := fmt.Sprintf("-p%s", os.Getenv("DB_PASS"))
	for _, query := range queries {
		err := exec.Command("mysql", "-h", os.Getenv("DB_HOST"), "-uroot", password, "gobm_t", "-e", query).Run()
		if err != nil {
			m.t.Fatalf("can not create: %v", err)
		}
	}
}

func (m testMysql) Truncates(tables []string) {
	password := fmt.Sprintf("-p%s", os.Getenv("DB_PASS"))
	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE %s", table)
		err := exec.Command("mysql", "-h", os.Getenv("DB_HOST"), "-uroot", password, "gobm_t", "-e", query).Run()
		if err != nil {
			m.t.Fatalf("can not truncate: %v", err)
		}
	}
}
