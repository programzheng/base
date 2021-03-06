package helper

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

func ConvertToString(any interface{}) string {
	switch value := any.(type) {
	case string:
		return value
	case int:
		return strconv.Itoa(value)
	case int64:
		return strconv.Itoa(int(value))
	case uint:
		return strconv.Itoa(int(value))
	default:
		log.Panic("ConvertToString error")
	}
	return ""
}

func ConvertInterfaceToIntMap(i []interface{}) []int {
	m := make([]int, len(i))
	for value := range i {
		m[value] = i[value].(int)
	}
	return m
}
