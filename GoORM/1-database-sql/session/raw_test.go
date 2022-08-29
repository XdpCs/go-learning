package session

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var TestDB *sql.DB

// go test 运行
func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "../fyy.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB)
}

func TestSession_Exec(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("drop table if exists user;").Exec()
	_, _ = s.Raw("create table user(Name text);").Exec()
	result, _ := s.Raw("insert into user(`Name`) values (?),(?)", "fyy", "xdp").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal("expect 2,but got", count)
	}
}

func TestSession_QueryRows(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("drop table if exists user;").Exec()
	_, _ = s.Raw("create table user(Name text);").Exec()
	row := s.Raw("select count(*) from user").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil || count != 0 {
		t.Fatal("failed to query db", err)
	}
}
