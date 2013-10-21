package configurator

import (
	"reflect"
)

func meldStructs(parent, target interface{}) {
	meldValueStructs(reflect.ValueOf(parent), reflect.ValueOf(target))
}

func meldValueStructs(parent, target reflect.Value) {
	if target.Kind() == reflect.Ptr {
		target = target.Elem()
	}
	if parent.Kind() == reflect.Ptr {
		parent = parent.Elem()
	}

	for i, n := 0, parent.NumField(); i < n; i++ {
		if target.Field(i).CanSet() {
			switch parent.Field(i).Kind() {
			case reflect.Struct:
				meldValueStructs(parent.Field(i), target.Field(i))
			case reflect.Array, reflect.Slice, reflect.String:
				if parent.Field(i).Len() > 0 {
					target.Field(i).Set(parent.Field(i))
				}
			default:
				if reflect.Zero(target.Field(i).Type()).Interface() != parent.Field(i).Interface() {
					target.Field(i).Set(parent.Field(i))
				}
			}
		}
	}
}
