package main

import (
	"GoMysql/config"
	schema "GoMysql/db"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbSelect *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", config.DBConfig.URL)
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	dbSelect = db
}
func main() {
	var person []schema.Person
	defer dbSelect.Close()
	err := dbSelect.Select(&person, "select * from person")
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	fmt.Println("select succ:", person)
}
