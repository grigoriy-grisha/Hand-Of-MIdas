package common

import (
	"math/rand"
	"reflect"
	"strconv"
)

func Rand20Percent() bool {
	return rand.Intn(100) < 20
}

func Typeof(value interface{}) string {
	return reflect.TypeOf(value).Name()
}

func KeyContains(key string, m map[string]interface{}) bool {
	_, exists := m[key]
	return exists
}

func SafeWriteToMap(m map[string]interface{}, key string, value interface{}) bool {
	if Typeof(m[key]) == "int" {
		convertedValue, err := strconv.Atoi(value.(string))

		if err != nil {
			return false
		}

		m[key] = convertedValue
	} else {
		m[key] = value
	}

	return true
}
