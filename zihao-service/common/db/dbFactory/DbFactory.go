package dbFactory

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/mysql"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlite"
	"github.com/zihao-boy/zihao/zihao-service/config"
)

const (
	Cache_sqlite = "sqlite"
	Cache_mysql  = "local"
)

func Init() {
	dbSwatch := config.G_AppConfig.Db

	if Cache_mysql == dbSwatch {
		mysql.InitGorm()
	}

	if Cache_sqlite == dbSwatch {
		sqlite.InitSqlite()
	}

}
