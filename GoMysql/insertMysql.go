package main

import (
	"GoMysql/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbInsert *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", config.DBConfig.URL)
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	dbInsert = database
}

func main() {
	r, err := dbInsert.Exec("insert into person(username,sex,email)values (?,?,?)", "fyy", "superman", "xdpcsyy@gmail.com")
	defer dbInsert.Close()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	fmt.Println("insert succ:", id)
}
