package module

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConDB() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "crud"
	//var db *sql.DB
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	//db, err = sql.Open("mysql", "didiyudha:ytrewq@/blog")
	if err != nil {
		panic(err.Error())
	} else {

		fmt.Println("Connected to DB!")
	}
	return db
}
