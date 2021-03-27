package main

import (
	"fmt"
	"reflect"
)

// 定义结构体 Response
type Response struct {
	Code    int    `json:"code" require:"true"`
	Message string `json:"message" require:"true"`
}

func main() {
	resp := Response{Code: 200, Message: "Success"}
	fmt.Println("Old: ", resp)
	var Iresp = &resp
	// 获取 Iresp 的 Value 对象
	obj := reflect.ValueOf(Iresp)
	v := obj.Elem()
	fmt.Print(v.FieldByName("Code"))
	// v.NumField() 为结构体字段数量
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			if v.Field(i).CanAddr() {
				v.Field(i).SetString("Update")
			}
		case reflect.Int:
			if v.Field(i).CanAddr() {
				v.Field(i).SetInt(210)
			}
		}
	}
	fmt.Println("Update: ", resp)
}
