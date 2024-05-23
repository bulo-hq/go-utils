package db

import (
	"database/sql"
	"fmt"
)

func GetSQLDB(
	host string,
	port string,
	dbname string,
	user string,
	password string,
	sslmode string,
	sslrootcert string,
	sslcert string,
	sslkey string,
) (*sql.DB, error) {
	uri := fmt.Sprintf("host=%s", host)
	uri += fmt.Sprintf(" port=%s", port)
	uri += fmt.Sprintf(" dbname=%s", dbname)
	uri += fmt.Sprintf(" user=%s", user)
	uri += fmt.Sprintf(" password=%s", password)
	uri += fmt.Sprintf(" sslmode=%s", sslmode)

	if sslrootcert != "" {
		uri += fmt.Sprintf(" sslrootcert=%s", sslrootcert)

		if sslcert != "" && sslkey != "" {
			uri += fmt.Sprintf(" sslcert=%s sslkey=%s", sslcert, sslkey)
		}
	}

	db, err := sql.Open("pgx", uri)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %s", err)
	}

	return db, err
}
