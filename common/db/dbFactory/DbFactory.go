package dbFactory

import (
	"github.com/zihao-boy/zihao/common/db/mysql"
	"github.com/zihao-boy/zihao/common/db/sqlite"
	"github.com/zihao-boy/zihao/config"
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
