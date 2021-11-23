package mysql

import (
	"github.com/kataras/golog"
	"github.com/zihao-boy/zihao/zihao-service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var G_DB *gorm.DB

func InitGorm() {
	var (
		err    error
		db     *gorm.DB
		config = config.G_DBConfig
		//root:root@tcp(localhost:3306)/bst?charset=utf8&parseTime=True&loc=Local
		url = config.DBConnUrl()
	)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       url,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		goto ERR
	}

	G_DB = db
	return
ERR:
	golog.Fatalf("~~> Mysql的gorm初始化错误,原因:%s", err.Error())
	return
}
