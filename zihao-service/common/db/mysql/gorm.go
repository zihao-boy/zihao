package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"github.com/zihao-boy/zihao/zihao-service/common/utils"
	"github.com/zihao-boy/zihao/zihao-service/config"
	"strconv"
	"strings"
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

/**
解析sql 引擎
		`select * count from host_group t
						where
						1=1
						$if name != '' then
						and t.name = #name#
						$endif
						$if name != '' then
						and t.name = #name#
						$endif
 */
func ParseSql(sql string,param map[string]interface{}) (string,[]interface{}){

	var (
		newSql string
		tmpSql string
		tmpWhere string //条件
		tmpWhereBool bool
		sqlParams []interface{}
	)

	sqls := strings.Split(sql,"$endif")

	for _,sql := range sqls{
		//不包含条件语句
		if sql == "" || !strings.Contains(sql,"$if"){
			newSql += sql
			continue
		}

		if !strings.HasPrefix(sql,"$if"){
			newSql += sql[0:strings.Index(sql,"$if")]
		}

		//条件语句
		tmpSql = sql[strings.Index(sql,"$if")+3:len(sql)]

		// 条件
		tmpWhere = tmpSql[0:strings.Index(tmpSql,"then")]

		tmpWhereBool = parseExpression(tmpWhere,param)

		if !tmpWhereBool {
			continue
		}

		tmpWhere = tmpSql[strings.Index(tmpSql,"then")+4:len(tmpSql)]
		newSql += tmpWhere

	}

	fmt.Print(newSql)

	//处理参数
	newSql,sqlParams = parseParam(newSql, param )


	return newSql,sqlParams
}

/**
解析参数
 */
func parseParam(sql string, param map[string]interface{}) (string,[]interface{}) {

	var (
		sqls []string
		currentSqlNew string
		sqlParams = make([]interface{},0)

	)

	sqls = strings.Split(sql,"#")

	for sqlIndex,tmpSql := range sqls{
		if sqlIndex % 2 == 0{
			currentSqlNew += tmpSql
			continue
		}
		currentSqlNew += "?"
		sqlParams = append(sqlParams,utils.ParseObjectValueFromInterface(param,strings.Trim(tmpSql," ")))
	}

	return currentSqlNew,sqlParams

}

/**
String[] sqls = currentSql.split("#");
            String currentSqlNew = "";
            for (int sqlIndex = 0; sqlIndex < sqls.length; sqlIndex++) {
                if (sqlIndex % 2 == 0) {
                    currentSqlNew += sqls[sqlIndex];
                    continue;
                }

currentSqlNew += "?";
Object param = params.getString(sqls[sqlIndex]);
if (params.get(sqls[sqlIndex]) instanceof Integer) {
param = params.getInteger(sqls[sqlIndex]);
}
//这里对 page 和 rows 特殊处理 ，目前没有想到其他的办法
if (StringUtils.isNumeric(param.toString()) && "page,rows,row".contains(sqls[sqlIndex])) {
param = Integer.parseInt(param.toString());
}
currentParams.add(param);
//currentSqlNew += params.get(sqls[sqlIndex]) instanceof Integer ? params.getInteger(sqls[sqlIndex]) : "'" + params.getString(sqls[sqlIndex]) + "'";
}
}
 */

/**
表达式 解析
 */
func parseExpression(tmpWhere string, param map[string]interface{}) bool {

	var (
		tmpWhereName string //条件name
		tmpWhereValue string //条件值
		paramValue string
	)

	if strings.Contains(tmpWhere,"!="){

		tmpWhereName = strings.Trim(tmpWhere[0:strings.Index(tmpWhere,"!=")]," ")
		tmpWhereValue = strings.Trim(tmpWhere[strings.Index(tmpWhere,"!=")+2:len(tmpWhere)]," ")
		tmpWhereValue = strings.ReplaceAll(tmpWhereValue,"'","")
		paramValue = utils.ParseStringValueFromInterface(param,tmpWhereName)
		if tmpWhereValue != paramValue{
			return true;
		}

	} else if strings.Contains(tmpWhere,"=="){
		tmpWhereName = strings.Trim(tmpWhere[0:strings.Index(tmpWhere,"==")]," ")
		tmpWhereValue = strings.Trim(tmpWhere[strings.Index(tmpWhere,"==")+2:len(tmpWhere)]," ")
		tmpWhereValue = strings.ReplaceAll(tmpWhereValue,"'","")
		if tmpWhereValue == utils.ParseStringValueFromInterface(param,tmpWhereName){
			return true;
		}
	} else if strings.Contains(tmpWhere,">"){
		tmpWhereName = strings.Trim(tmpWhere[0:strings.Index(tmpWhere,">")]," ")
		tmpWhereValue = strings.Trim(tmpWhere[strings.Index(tmpWhere,">")+1:len(tmpWhere)]," ")
		tmpWhereValue = strings.ReplaceAll(tmpWhereValue,"'","")
		if tmpWhereValueInt,_ := strconv.ParseInt(tmpWhereValue,10,64);
			tmpWhereValueInt > utils.ParseIntValueFromInterface(param,tmpWhereName){
			return true;
		}
	}else if strings.Contains(tmpWhere,"<"){
		tmpWhereName = strings.Trim(tmpWhere[0:strings.Index(tmpWhere,"<")]," ")
		tmpWhereValue = strings.Trim(tmpWhere[strings.Index(tmpWhere,"<")+1:len(tmpWhere)]," ")
		tmpWhereValue = strings.ReplaceAll(tmpWhereValue,"'","")
		if tmpWhereValueInt,_ := strconv.ParseInt(tmpWhereValue,10,64);
			tmpWhereValueInt < utils.ParseIntValueFromInterface(param,tmpWhereName){
			return true;
		}
	}else if strings.Contains(tmpWhere,">="){
		tmpWhereName = strings.Trim(tmpWhere[0:strings.Index(tmpWhere,">=")]," ")
		tmpWhereValue = strings.Trim(tmpWhere[strings.Index(tmpWhere,">=")+2:len(tmpWhere)]," ")
		tmpWhereValue = strings.ReplaceAll(tmpWhereValue,"'","")
		if tmpWhereValueInt,_ := strconv.ParseInt(tmpWhereValue,10,64);
			tmpWhereValueInt >= utils.ParseIntValueFromInterface(param,tmpWhereName){
			return true;
		}
	}else if strings.Contains(tmpWhere,"<="){
		tmpWhereName = strings.Trim(tmpWhere[0:strings.Index(tmpWhere,"<=")]," ")
		tmpWhereValue = strings.Trim(tmpWhere[strings.Index(tmpWhere,"<=")+2:len(tmpWhere)]," ")
		tmpWhereValue = strings.ReplaceAll(tmpWhereValue,"'","")
		if tmpWhereValueInt,_ := strconv.ParseInt(tmpWhereValue,10,64);
			tmpWhereValueInt <= utils.ParseIntValueFromInterface(param,tmpWhereName){
			return true;
		}
	}

	return false;
}

/**
for (String oSql : oSqls) {
            logger.debug("处理if 节点，当前处理的oSql=" + oSql + "总的oSqls = " + oSqls);
            if (StringUtil.isNullOrNone(oSql) || !oSql.contains("<if")) {
                newSql.append(oSql);
                continue;
            }
            if (!oSql.startsWith("<if")) {
                newSql.append(oSql.substring(0, oSql.indexOf("<if")));
            }

            tmpSql = oSql.substring(oSql.indexOf("<if")) + "</if>";

            Element root = DocumentHelper.parseText(tmpSql).getRootElement();

            String condition = root.attribute("test").getValue();

            Object condObj = Ognl.parseExpression(condition);

            Object value = Ognl.getValue(condObj, requestParams);

            if (value instanceof Boolean) {
                conditionResult = (Boolean) value;
            } else {
                throw new BusinessException(ResponseConstant.RESULT_CODE_INNER_ERROR, "配置错误，if语句配置错误 " + condition);
            }

            if (conditionResult) {
                newSql.append(root.getText());
            }


        }
 */

/**
封装 查询列表
 */
func SelectList(sql string,param map[string]interface{},callback func(db *gorm.DB)) {
	var(
		newSql string
		sqlParams []interface{}
	)
	newSql,sqlParams = ParseSql(sql,param)
	db := G_DB.Raw(newSql,sqlParams...)
	callback(db)
	//if err:=db.Scan(&out).Error; err !=nil{
	//	return nil,err
	//}

}

/**
封装 查询单个
*/
func SelectOne(sql string,param map[string]interface{},callback func(db *gorm.DB)) {
	var(
		newSql string
		sqlParams []interface{}
	)
	newSql,sqlParams = ParseSql(sql,param)
	db := G_DB.Raw(newSql,sqlParams...)
	callback(db)
}