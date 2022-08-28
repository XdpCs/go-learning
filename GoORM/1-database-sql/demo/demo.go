package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, _ := sql.Open("sqlite3", "fyy.db")
	defer func() {
		_ = db.Close()
	}()
	_, _ = db.Exec("drop table if exists user;")
	_, _ = db.Exec("create table user(Name text);")
	result, err := db.Exec("insert into user('Name') values (?),(?)", "fyy", "xdp")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	row := db.QueryRow("select name from user limit 1")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}
