package helper

import "reflect"

func GetStructName(s interface{}) string {
	return reflect.TypeOf(&s).Elem().Name()
}
