package sqlTemplate

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/mysql"
	"github.com/zihao-boy/zihao/zihao-service/common/utils"
	"strconv"
	"strings"
)

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

	//update 时 如果 where 前面有, 则删除

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
func SelectList(sql string,param map[string]interface{},callback func(db *gorm.DB),cacheSql bool) {
	var(
		newSql string
		sqlParams []interface{}
	)
	if cacheSql{
		serviceSql := utils.GetServiceSql(sql)
		sql = serviceSql.SqlText
	}
	newSql,sqlParams = ParseSql(sql,param)
	db := mysql.G_DB.Raw(newSql,sqlParams...)
	callback(db)
	//if err:=db.Scan(&out).Error; err !=nil{
	//	return nil,err
	//}

}

/**
封装 查询单个
*/
func SelectOne(sql string,param map[string]interface{},callback func(db *gorm.DB),cacheSql bool) {
	var(
		newSql string
		sqlParams []interface{}
	)
	if cacheSql{
		serviceSql := utils.GetServiceSql(sql)
		sql = serviceSql.SqlText
	}
	newSql,sqlParams = ParseSql(sql,param)
	db := mysql.G_DB.Raw(newSql,sqlParams...)
	callback(db)
}


/**
封装 新增数据
*/
func Insert(sql string,param map[string]interface{},cacheSql bool) error {
	var(
		newSql string
		sqlParams []interface{}
	)
	if cacheSql{
		serviceSql := utils.GetServiceSql(sql)
		sql = serviceSql.SqlText
	}
	if sql == ""{
		return errors.New("sql 配置错误 值为空")
	}
	newSql,sqlParams = ParseSql(sql,param)
	db := mysql.G_DB.Exec(newSql,sqlParams...)

	return db.Error
}


/**
封装 修改数据
*/
func Update(sql string,param map[string]interface{},cacheSql bool) error {
	var(
		newSql string
		sqlParams []interface{}
	)
	if cacheSql{
		serviceSql := utils.GetServiceSql(sql)
		sql = serviceSql.SqlText
	}
	newSql,sqlParams = ParseSql(sql,param)
	db := mysql.G_DB.Exec(newSql,sqlParams...)
	return db.Error
}

/**
封装  删除数据
*/
func Delete(sql string,param map[string]interface{},cacheSql bool) error{
	var(
		newSql string
		sqlParams []interface{}
	)
	if cacheSql{
		serviceSql := utils.GetServiceSql(sql)
		sql = serviceSql.SqlText
	}
	newSql,sqlParams = ParseSql(sql,param)
	db := mysql.G_DB.Exec(newSql,sqlParams...)
	return db.Error
}
