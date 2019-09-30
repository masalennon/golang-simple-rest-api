package data

import (
	"log"
	"database/sql"

)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}