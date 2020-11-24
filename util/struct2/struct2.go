package struct2

import (
	"encoding/json"
	"reflect"
)

func StructToEndMap(stru interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(stru)
	json.Unmarshal(j, &m)
	return m
}
func GetStructName(info interface{}) string {
	return reflect.TypeOf(info).Name()
}
