package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var jsonData map[string]interface{}
var DBConfig dbConfig

type dbConfig struct {
	Database string
	User     string
	Password string
	Host     string
	Port     string
	URL      string
}

func initJson() {
	bytes, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		fmt.Printf("ReadFile:%v\n", err)
		os.Exit(-1)
	}
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("json parse fail", err.Error())
		os.Exit(-1)
	}
}

func initDB() {
	DBConfig.Host = jsonData["Host"].(string)
	DBConfig.Database = jsonData["DataBase"].(string)
	DBConfig.User = jsonData["User"].(string)
	DBConfig.Password = jsonData["Password"].(string)
	DBConfig.Port = jsonData["Port"].(string)
	// url例子：root:root@tcp(127.0.0.1:3306)/test
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database)
	DBConfig.URL = url
}

func init() {
	initJson()
	initDB()
}
