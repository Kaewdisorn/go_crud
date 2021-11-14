package module

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConDB() (db *sql.DB) {

	//fmt.Println("ConDB Opened")
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "crud"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	} else {

		fmt.Println("Connected to DB!")
	}

	return db
}
