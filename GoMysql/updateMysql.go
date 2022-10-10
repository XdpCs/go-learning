package main

import (
	"GoMysql/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbUpdate *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", config.DBConfig.URL)
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	dbUpdate = db
}

func main() {
	result, err := dbUpdate.Exec("update person set username=? where user_id=?", "fyyxdp", 2)
	defer dbUpdate.Close()
	if err != nil {
		fmt.Println("exec failed,", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("rows failed,", err)
	}
	fmt.Println("update succ:", rowsAffected)
}
