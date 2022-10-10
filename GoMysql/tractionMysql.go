package main

import (
	"GoMysql/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// mysql事务特性：
// 1.原子性
// 2.一致性
// 3.隔离性
// 4.持久性
// Begin() 开始事务
// Commit() 提交事务
// Rollback() 回滚事务

var dbTraction *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", config.DBConfig.URL)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	dbTraction = db
}

func main() {
	conn, err := dbTraction.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}
	result, err := conn.Exec("insert into person(username,sex,email)values (?,?,?)", "xdp", "man", "xdpfyy@qq.com")
	if err != nil {
		fmt.Println("exec failed,", err)
		conn.Rollback()
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ", id)
	result, err = conn.Exec("insert into person(username,sex,email)values (?,?,?)", "fyy", "man", "xdpfyy@qq.com")
	if err != nil {
		fmt.Println("exec failed,", err)
		conn.Rollback()
		return
	}
	id, err = result.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)
	conn.Commit()
}
