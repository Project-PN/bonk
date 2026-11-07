package model

import (
	"fmt"
	"reflect"
)

// grab the fields from a model struct
func inspect(t reflect.Type) ([]reflect.StructField, error) {
	seen := make(map[reflect.Type]bool)
	for t != nil && t.Kind() == reflect.Ptr {
		if seen[t] {
			return nil, fmt.Errorf("circular reference detected for type %s", t)
		}

		seen[t] = true
		t = t.Elem()
	}
	if t == nil || t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("model must be a struct, got %v", t)
	}

	// visible fields
	var fields []reflect.StructField
	for _, field := range reflect.VisibleFields(t) {
		if field.IsExported() {
			fields = append(fields, field)
		}
	}

	return fields, nil
}
