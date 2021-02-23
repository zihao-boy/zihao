package utils

import (
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"strconv"
)

/**
从interface 对象中反射 出值
 */
func ParseIntValueFromInterface(param map[string]interface{},name string) int64 {
	if name == "Page" {
		return param["PageDto"].(dto.PageDto).Page
	}
	if name == "Row" {
		return param["PageDto"].(dto.PageDto).Row
	}
	return param[name].(int64)
}

func ParseStringValueFromInterface(param map[string]interface{},name string) string {
	if name == "Page" {
		return strconv.FormatInt(param["PageDto"].(dto.PageDto).Page,10)
	}
	if name == "Row" {
		return strconv.FormatInt(param["PageDto"].(dto.PageDto).Row,10)
	}
	return param[name].(string)

}

func ParseObjectValueFromInterface(param map[string]interface{},name string) interface{} {
	if name == "Page" {
		return param["PageDto"].(dto.PageDto).Page
	}
	if name == "Row" {
		return param["PageDto"].(dto.PageDto).Row
	}
	return param[name]
}


