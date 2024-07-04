package db

import (
	"database/sql"
	"fmt"
)

func GetSQLDB(schemaName string) (*sql.DB, error) {
	db, err := sql.Open("pgx", getURI(schemaName))
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %s", err)
	}

	return db, err
}
