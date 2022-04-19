package helper

func MergeStringInterfaceMap(map1 map[string]interface{}, map2 map[string]interface{}) map[string]interface{} {
	for k, v := range map2 {
		map1[k] = v
	}
	return map1
}
