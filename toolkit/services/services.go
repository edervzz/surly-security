package services

import (
	"fmt"
	"reflect"
)

var serviceCollection map[string]any = make(map[string]any)

func AddSingleton[T interface{}](instance T) {
	// create empty object and get info
	empty := (*T)(nil)
	t := reflect.TypeOf(empty)
	name := fmt.Sprint(t.Elem())
	// save instance
	serviceCollection[name] = instance
}

func Get[T interface{}]() T {
	// create empty object and get info
	empty := (*T)(nil)
	t := reflect.TypeOf(empty)
	name := fmt.Sprint(t.Elem())
	// get instance
	return serviceCollection[name].(T)
}
