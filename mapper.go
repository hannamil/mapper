package mapper

import (
	"errors"
	"reflect"
)

var (
	ZeroValue reflect.Value
)

func init() {
	ZeroValue = reflect.Value{}
}

func Mapper(source, dest interface{}) (interface{}, error) {
	sourceElem := reflect.ValueOf(source)
	destElem := reflect.ValueOf(dest)
	if sourceElem == ZeroValue {
		return nil, errors.New("source object is not legal value")
	}
	if destElem == ZeroValue {
		return nil, errors.New("dest object is not legal value")
	}

	sourceKind := reflect.ValueOf(source).Type().Kind()
	destKind := reflect.ValueOf(dest).Type().Kind()
	if sourceKind != destKind {
		return nil, errors.New("type error")
	}

	return structMapper(sourceElem, destElem, destKind)
}

func structMapper(source, dest reflect.Value, kind reflect.Kind) (interface{}, error) {
	if kind == reflect.Ptr {
		return elementToStruct(source.Elem(), dest.Elem())
	}
	return elementToStruct(source, reflect.New(dest.Type()).Elem())
}

func elementToStruct(source, dest reflect.Value) (interface{}, error) {
	for i := 0; i < source.NumField(); i++ {
		name := source.Type().Field(i).Name
		destField := dest.FieldByName(name)
		if destField.IsValid() {
			destField.Set(source.Field(i))
		}
	}
	return dest, nil
}
