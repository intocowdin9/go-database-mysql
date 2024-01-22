package go_database_mysql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	dataSource := "root:@tcp(localhost:3306)/go_database?parseTime=true"
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	// gunakan DB
}
