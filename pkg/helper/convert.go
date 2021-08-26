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
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.Itoa(int(value))
	case float64:
		return strconv.FormatFloat(value, 'f', 2, 64)
	default:
		log.Panic("ConvertToString error")
	}
	return ""
}

func ConvertToInt(any interface{}) int {
	switch value := any.(type) {
	case int:
		return value
	case float64:
		return int(value)
	case string:
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Panic("convert string to int error:", err)
		}
		return i
	default:
		log.Panic("ConvertToInt error")
	}
	return -1
}

func ConvertToFloat64(any interface{}) float64 {
	switch value := any.(type) {
	case int:
		return float64(value)
	case float64:
		return value
	case string:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Panic("convert string to float64  error")
		}
		return f
	default:
		log.Panic("ConvertToFloat64 error")
	}
	return -1.00
}

func ConvertToBool(any interface{}) bool {
	switch value := any.(type) {
	case string:
		b, err := strconv.ParseBool(value)
		if err != nil {
			log.Panic("convert string to bool error")
		}
		return b
	}
	return false
}

func ConvertInterfaceToIntMap(i []interface{}) []int {
	m := make([]int, len(i))
	for value := range i {
		m[value] = i[value].(int)
	}
	return m
}
