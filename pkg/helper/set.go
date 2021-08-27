package helper

import (
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"
)

func SliceToSet(any interface{}) interface{} {
	rv := reflect.ValueOf(any)
	//check is Ptr
	rv = reflect.Indirect(rv)
	if rv.Kind() != reflect.Slice {
		log.Panic("helper SliceToSet can't use except for the type of slice")
	}
	if rv.IsNil() {
		return nil
	}
	set := make(map[interface{}]struct{}, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		s := rv.Index(i).Interface()
		switch s.(type) {
		case string:
			set[s.(string)] = struct{}{}
		}
	}
	setKeys := make([]interface{}, 0, len(set))
	for k := range set {
		fmt.Println(k)
		setKeys = append(setKeys, k)
	}
	return setKeys
}
