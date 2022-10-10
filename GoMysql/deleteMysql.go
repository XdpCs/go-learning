package main

import (
	"GoMysql/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbDelete *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", config.DBConfig.URL)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	dbDelete = db
}

func main() {
	res, err := dbDelete.Exec("delete from person where user_id = ?", 2)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed,", err)
	}
	fmt.Println("delete succ:", rowsAffected)
}
