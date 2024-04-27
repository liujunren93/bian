package utils

import (
	"fmt"
	"reflect"
)

func Struct2MapNoZero(src any) (map[string]any, error) {
	if v, ok := src.(map[string]any); ok {
		return v, nil
	}
	t := reflect.TypeOf(src)
	v := reflect.ValueOf(src)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {

		return nil, fmt.Errorf("src is not a struct")
	}
	fieldNum := t.NumField()
	var res = make(map[string]any, fieldNum)
	for i := 0; i < fieldNum; i++ {
		if t.Field(i).Name[0] >= 97 {
			continue
		}
		if t.Field(i).Tag.Get("json") == "-" {
			continue
		}

		if v.Kind() == reflect.Ptr {
			field := v.Elem().Field(i)
			if !field.IsZero() {
				res[t.Field(i).Tag.Get("json")] = field.Interface()
			}
		} else {
			field := v.Field(i)
			if !field.IsZero() {
				res[t.Field(i).Tag.Get("json")] = field.Interface()
			}
		}

	}
	return res, nil
}
