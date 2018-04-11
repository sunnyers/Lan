package lan

//两个map合并，如果有相同的key 则后面map2的值替换掉map1的值
func MapMerge(map1 map[string]interface{},map2 map[string]interface{}) map[string]interface{}{
	for k,v := range map2{
		map1[k] = v
	}
	return map1
}

func Isset(data map[string]interface{},key string) (string,bool) {
	if val,ok := data[key]; ok {
		return val.(string),true
	}
	return "",false
}
