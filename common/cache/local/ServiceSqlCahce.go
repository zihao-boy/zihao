package local

import (
	"encoding/json"
	"fmt"

	"github.com/zihao-boy/zihao/entity/dto/serviceSql"
	"github.com/zihao-boy/zihao/system/mapper"
)

func (r *Local) SaveServiceSql(serviceSqlDto serviceSql.ServiceSqlDto) (err error) {
	data, err := json.Marshal(serviceSqlDto)
	r.client.Set(serviceSqlDto.SqlCode, data,
		-1)
	return err
}

func (r *Local) GetServiceSql(sqlCode string) (serviceSql serviceSql.ServiceSqlDto, err error) {
	data, _ := r.client.Get(sqlCode)
	json.Unmarshal(data.([]byte), &serviceSql)
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

	fmt.Print("123123", serviceSqlDtos)

	if err != nil {
		fmt.Print("加载sql 失败", err)
		return
	}

	for _, item := range serviceSqlDtos {
		G_Local.SaveServiceSql(*item)
	}
}
