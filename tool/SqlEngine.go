package tool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func GetDb() *sql.DB {
	return Db
}

func SqlEngine() {
	db, err := sql.Open("mysql", "root:hjt82572799@tcp(localhost:3306)/gin-message-board")
	if err != nil {
		fmt.Println(err)
		return
	}
	Db = db
}
