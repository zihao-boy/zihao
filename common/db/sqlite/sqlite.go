package sqlite

import (
	"github.com/zihao-boy/zihao/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var S_DB *gorm.DB

/**

初始化 数据库
**/
func InitSqlite() {
	//db, err := sql.Open("sqlite3", "./db/zihao_dev.db")
	sqlitePath := config.G_AppConfig.SqlitePath
	if len(sqlitePath) == 0{
		sqlitePath = "./db/zihao.db"
	}
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	checkErr(err)
	S_DB = db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
