package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var S_DB *gorm.DB

/**

初始化 数据库
**/
func InitSqlite() {
	//db, err := sql.Open("sqlite3", "./db/zihao.db")
	db, err := gorm.Open(sqlite.Open("./db/zihao.db"), &gorm.Config{})
	checkErr(err)
	S_DB = db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
