package reflection

import (
	"errors"
	"reflect"
	"strings"
)

const defaultSeparator = ","

var empty interface{}

var ErrFiledValueIsNotString = errors.New("the value for the requested field is not a string")

var ErrTypeNotStructure = errors.New("type not structure")

var ErrFieldWithRequestedTagNotExist = errors.New("field with requested tag not exist")

var ErrInterfaceNotFound = errors.New("no nil interface does not exist")

func StringFieldByName(i interface{}, field string) (string, error) {
	v := reflect.Indirect(reflect.ValueOf(i)).FieldByName(field)

	if v.Kind() == reflect.String {
		return v.String(), nil
	}

	return "", ErrFiledValueIsNotString
}

func FirstNotNilInterface(str interface{}) (interface{}, error) {
	v := reflect.ValueOf(str)

	if v.Kind() != reflect.Struct {
		return empty, ErrTypeNotStructure
	}

	for i := 0; i < v.NumField(); i++ {
		if field := v.Field(i); field.IsValid() &&
			(field.Kind() == reflect.Ptr ||
				field.Kind() == reflect.Interface) &&
			!field.IsNil() {
			return field.Interface(), nil
		}
	}

	return empty, ErrInterfaceNotFound
}

func FieldByTag(str interface{}, tagName, tagValue string) (reflect.Value, error) {
	v := reflect.ValueOf(str)
	t := reflect.TypeOf(str)

	if v.Kind() != reflect.Struct {
		return reflect.Value{}, ErrTypeNotStructure
	}

	for i := 0; i < v.NumField(); i++ {
		if tValue, ok := t.Field(i).Tag.Lookup(tagName); ok {
			for _, s := range strings.Split(tValue, defaultSeparator) {
				if s == tagValue {
					return v.Field(i), nil
				}
			}
		}
	}

	return reflect.Value{}, ErrFieldWithRequestedTagNotExist
}
