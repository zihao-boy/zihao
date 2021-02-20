package utils

/**
从interface 对象中反射 出值
 */
func ParseIntValueFromInterface(param map[string]interface{},name string) int64 {
	return param[name].(int64)
}

func ParseStringValueFromInterface(param map[string]interface{},name string) string {
	return param[name].(string)

}

func ParseObjectValueFromInterface(param map[string]interface{},name string) interface{} {
	return param[name]
}


