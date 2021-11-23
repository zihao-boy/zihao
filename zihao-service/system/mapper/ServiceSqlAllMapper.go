package mapper

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/mysql"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlite"
	"github.com/zihao-boy/zihao/zihao-service/config"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/serviceSql"
)

const (
	Cache_sqlite = "sqlite"
	Cache_mysql  = "local"
)

type ServiceSqlAllMapper struct {
}

/**
查询用户


*/
func (*ServiceSqlAllMapper) GetServiceSqls(serviceSqlDto serviceSql.ServiceSqlDto) ([]*serviceSql.ServiceSqlDto, error) {
	var serviceSqlDtos []*serviceSql.ServiceSqlDto

	dbSwatch := config.G_AppConfig.Db

	if Cache_mysql == dbSwatch {
		db := mysql.G_DB.Raw("select * from service_sql")
		if err := db.Scan(&serviceSqlDtos).Error; err != nil {
			return nil, err
		}
	}

	if Cache_sqlite == dbSwatch {
		db, _ := sqlite.S_DB.Query("select * from service_sql")
		if err := db.Scan(&serviceSqlDtos); err != nil {
			return nil, err
		}
	}
	return serviceSqlDtos, nil
}
