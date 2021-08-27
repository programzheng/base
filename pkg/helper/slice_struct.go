package helper

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

func Pluck(any interface{}, key string) interface{} {
	//interface to reflect.Value
	rv := reflect.ValueOf(any)
	//check is Ptr
	rv = reflect.Indirect(rv)
	if rv.Kind() != reflect.Slice {
		log.Panic("helper pluck can't use except for the type of slice")
	}
	if rv.IsNil() {
		return nil
	}
	//use interface array
	out := make([]interface{}, rv.Len())
	//get target all type
	targetType := reflect.Invalid
	for i := 0; i < rv.Len(); i++ {
		s := rv.Index(i)
		//get target key's value
		target := s.FieldByName(key)
		//target type switch
		switch kind := target.Kind(); kind {
		case reflect.String:
			out[i] = target.String()
			targetType = targetType + reflect.String
		case reflect.Int64:
			out[i] = target.Int()
			targetType = targetType + reflect.Int
		}
	}
	//vail all target type
	switch {
	case targetType == reflect.String*reflect.Kind(rv.Len()):
		stringOut := make([]string, rv.Len())
		for i, v := range out {
			stringOut[i] = v.(string)
		}
		return stringOut
	case targetType == reflect.Int*reflect.Kind(rv.Len()):
		intOut := make([]int64, rv.Len())
		for i, v := range out {
			intOut[i] = v.(int64)
		}
		return intOut
	}
	return out
}
