package dbFactory

import (
	"fmt"
	mysqlUtil "github.com/zihao-boy/zihao/common/db/mysql"
	"github.com/zihao-boy/zihao/common/db/sqlite"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/dbLink"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"sync"
)

const (
	Cache_sqlite = "sqlite"
	Cache_mysql  = "local"
)

var (
	dbLinkDtos = make(map[string]*gorm.DB)
	lock       sync.Mutex
)

func Init() {
	dbSwatch := config.G_AppConfig.Db

	if Cache_mysql == dbSwatch {
		mysqlUtil.InitGorm()
	}

	if Cache_sqlite == dbSwatch {
		sqlite.InitSqlite()
	}
}

// exec sql

func ExecSql(dblinkDto dbLink.DbLinkDto, dbSqlDto dbLink.DbSqlDto) interface{} {

	var (
		datas []map[string]interface{}
	)
	db, err := initDbLink(dblinkDto)

	if err != nil {
		return result.Error(err.Error())
	}

	sql := strings.ToLower(dbSqlDto.Sql)

	sql = strings.ReplaceAll(sql, " ", "")
	sql = strings.ReplaceAll(sql, "/r", "")
	sql = strings.ReplaceAll(sql, "/n", "")

	if strings.HasPrefix(sql, "select") {
		rows, _ := db.Exec(dbSqlDto.Sql).Rows()
		cols, _ := rows.Columns()
		for rows.Next() {
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i, _ := range columns {
				columnPointers[i] = &columns[i]
			}
			// Scan the result into the column pointers...
			if err := rows.Scan(columnPointers...); err != nil {
				return err
			}
			// Create our map, and retrieve the value for each column from the pointers slice,
			// storing it in the map with the name of the column as the key.
			m := make(map[string]interface{})
			for i, colName := range cols {
				val := columnPointers[i].(*interface{})
				m[colName] = *val
			}
			// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
			fmt.Print(m)
			datas = append(datas, m)
		}
		return result.SuccessData(datas)
	} else {
		err := db.Exec(dbSqlDto.Sql).Error
		if err != nil {
			return result.Error(err.Error())
		}
		return result.Success()
	}

}

func initDbLink(dblinkDto dbLink.DbLinkDto) (*gorm.DB, error) {
	db := dbLinkDtos[dblinkDto.Id]
	if db != nil {
		return db, nil
	}
	lock.Lock()
	defer func() {
		lock.Unlock()
	}()
	// judge again
	db = dbLinkDtos[dblinkDto.Id]
	if db != nil {
		return db, nil
	}
	var (
		err error
		//root:root@tcp(localhost:3306)/bst?charset=utf8&parseTime=True&loc=Local
		//hc_things:wuxw2015@tcp(106.52.221.206:%!d(string=3306))/tt?charset=utf8&parseTime=True&loc=Local
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", dblinkDto.Username, dblinkDto.Password, dblinkDto.Ip, dblinkDto.Port, dblinkDto.DbName, "utf8")
	)

	fmt.Print(url)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       url,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	dbLinkDtos[dblinkDto.Id] = db
	return db, nil
}
