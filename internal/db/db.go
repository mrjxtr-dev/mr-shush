package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbSource string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbSource)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return db, nil
}
