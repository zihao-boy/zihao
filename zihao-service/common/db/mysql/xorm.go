package mysql

import (
_ "github.com/go-sql-driver/mysql"
"github.com/go-xorm/xorm"
"github.com/kataras/golog"
"github.com/zihao-boy/zihao/zihao-service/config"
"xorm.io/core"
)

var G_Xorm *xorm.Engine

// http://gobook.io/read/github.com/go-xorm/manual-en-US/chapter-02/1.mapping.html
func InitXorm() {
	var (
		err  error
		engine *xorm.Engine
		config = config.G_DBConfig
	)
	if engine, err = xorm.NewEngine(config.Mysql.Dialect, config.DBConnUrl()); err != nil {
		goto ERR
	}
	if err = engine.Ping(); err != nil {
		goto ERR
	}
	engine.ShowSQL(config.Mysql.ShowSql)
	engine.SetMapper(core.GonicMapper{})
	engine.SetMaxOpenConns(config.Mysql.MaxOpenConns)
	engine.SetMaxIdleConns(config.Mysql.MaxIdleConns)
	//engine.SetTZLocation(utils.SysTimeLocation)
	// 性能优化的时候才考虑，加上本机的SQL缓存
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//engine.SetDefaultCacher(cacher)
	G_Xorm = engine
	return
ERR:
	golog.Fatalf("~~> Mysql初始化错误,原因:%s", err.Error())
	return
}
