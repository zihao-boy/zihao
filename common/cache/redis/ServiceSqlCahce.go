package redis

import (
	"encoding/json"
	"fmt"

	"github.com/zihao-boy/zihao/entity/dto/serviceSql"
	"github.com/zihao-boy/zihao/system/mapper"
)

func (r *Redis) SaveServiceSql(serviceSqlDto serviceSql.ServiceSqlDto) (err error) {
	data, err := json.Marshal(serviceSqlDto)
	err = r.client.Set(serviceSqlDto.SqlCode, data,
		0).Err()
	return err
}

func (r *Redis) GetServiceSql(sqlCode string) (serviceSql serviceSql.ServiceSqlDto, err error) {
	data, err := r.client.Get(sqlCode).Result()
	json.Unmarshal([]byte(data), &serviceSql)
	return serviceSql, err
}

// Init
func InitServiceSql() {
	var (
		serviceSqlAllMapper mapper.ServiceSqlAllMapper
		serviceSqlDto       serviceSql.ServiceSqlDto = serviceSql.ServiceSqlDto{}
		serviceSqlDtos      []*serviceSql.ServiceSqlDto
		err                 error
	)

	serviceSqlDtos, err = serviceSqlAllMapper.GetServiceSqls(serviceSqlDto)

	if err != nil {
		fmt.Print("加载sql 失败", err)
		return
	}

	for _, item := range serviceSqlDtos {
		G_Redis.SaveServiceSql(*item)
	}
}
