package mysql


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"github.com/zihao-boy/zihao/zihao-service/config"
)

var G_DB *gorm.DB

func InitGorm() {
	var (
		err  error
		db   *gorm.DB
		config = config.G_DBConfig
		//root:root@tcp(localhost:3306)/bst?charset=utf8&parseTime=True&loc=Local
		url = config.DBConnUrl()
	)
	if db, err = gorm.Open(config.Mysql.Dialect, url); err != nil {
		goto ERR
	}
	if err = db.DB().Ping(); err != nil {
		goto ERR
	}
	//一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	db.SingularTable(true)
	db.LogMode(config.Mysql.ShowSql)
	db.DB().SetMaxOpenConns(config.Mysql.MaxOpenConns)
	db.DB().SetMaxIdleConns(config.Mysql.MaxIdleConns)
	G_DB = db
	return
ERR:
	golog.Fatalf("~~> Mysql的gorm初始化错误,原因:%s", err.Error())
	return
}
