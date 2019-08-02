package mycore

import (
	"reflect"
	"strings"
)

//Struct2Map
func Struct2Map(obj interface{}) map[string]interface{} {
	 t := reflect.TypeOf(obj)
	 v := reflect.ValueOf(obj)

	 var data = make(map[string]interface{})
	 for i := 0; i < t.NumField(); i++ {
	 	data[strings.ToUpper(t.Field(i).Name)] = v.Field(i).Interface()
	 }
	 return data
}
