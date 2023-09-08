package dbUtils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, openErr := sql.Open("sqlite3", "urls.db")

	if openErr != nil {
		return nil, openErr
	}

	initDbQuery, initErr := db.Prepare(`CREATE TABLE IF NOT EXISTS URLS(ID INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, URL TEXT NOT NULL)`)

	if initErr != nil {
		return nil, initErr
	}

	_, execErr := initDbQuery.Exec()
	if execErr != nil {
		return nil, execErr
	}

	return db, nil
}
