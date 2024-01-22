package go_database_mysql

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	dataSource := "root:@tcp(localhost:3306)/go_database?parseTime=true"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
