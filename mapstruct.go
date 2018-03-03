package mapstruct

import (
	"fmt"
	"github.com/serenize/snaker"
	"reflect"
)

func MapInterfaceToStruct(data map[string]interface{}, stru interface{}) error {

	structValue := reflect.ValueOf(stru).Elem()

	for n, v := range data {

		n = snaker.SnakeToCamel(n)

		structFieldValue := structValue.FieldByName(n)

		if !structFieldValue.IsValid() {
			continue
		}

		if structFieldValue.Type() != reflect.TypeOf(v) {

			return fmt.Errorf("value types did not match for field %s, expected %s, got %s", n, structFieldValue.Type().String(), reflect.TypeOf(v).String())

		}

		if !structFieldValue.CanSet() {
			return fmt.Errorf("struct field not settable")
		}

		structFieldValue.Set(reflect.ValueOf(v))

	}

	return nil

}

func StructToMapInterface(stru interface{}) (map[string]interface{}, error) {

	structType := reflect.TypeOf(stru).Elem()

	structValue := reflect.ValueOf(stru).Elem()

	ret := map[string]interface{}{}

	for i := 0; i < structType.NumField(); i++ {
		structField := structType.Field(i)

		structFieldName := structField.Name

		structFieldValue := structValue.FieldByName(structFieldName).Interface()

		if structFieldValue != nil {
			ret[snaker.CamelToSnake(structFieldName)] = structFieldValue
		}

	}

	return ret, nil
}

func MergeMaps(left, right map[string]interface{}) map[string]interface{} {

	for k, v := range right {
		left[k] = v
	}

	return left
}
