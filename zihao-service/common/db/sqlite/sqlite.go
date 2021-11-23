package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var S_DB *sql.DB

/**

初始化 数据库
**/
func InitSqlite() {
	db, err := sql.Open("sqlite3", "./db/zihao.db")
	checkErr(err)
	S_DB = db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
