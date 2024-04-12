package db

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func NewDB(dbURL, dbToken string) (*sql.DB, error) {
	url := fmt.Sprintf("%v?authToken=%v", dbURL, dbToken)
	db, err := sql.Open("libsql", url)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, err
}
