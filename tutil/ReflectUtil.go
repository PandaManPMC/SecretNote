package tutil

import (
	"fmt"
	"reflect"
)

type reflectUtil struct{}

var reflectUtilInstance *reflectUtil

func GetInstanceByReflectUtil() *reflectUtil {
	return reflectUtilInstance
}

func (*reflectUtil) ErrToMap(err error) map[string]string {
	if nil == err {
		return nil
	}
	m := make(map[string]string)
	re := reflect.ValueOf(err)
	rt := reflect.TypeOf(err)
	num := re.Elem().NumField()
	for i := 0; i < num; i++ {
		r := rt.Elem().Field(i)
		f := re.Elem().Field(i)
		m[r.Name] = fmt.Sprintf("%v", f)
	}
	return m
}
