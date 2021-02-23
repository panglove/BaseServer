package struct2

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

func GetStructKeyList(stru interface{})  []string {
	newMap := StructToEndMap(stru)

	var list []string

	for k:=range newMap {

		list = append(list,k)
	}
	return list

}
func GetStructSqlValueList(stru interface{})  []string {
	newMap := StructToEndMap(stru)

	var list []string

	for k:=range newMap {

		list = append(list,InterfaceStrSqlval(newMap[k]))
	}
	return list

}
func GetStructSqlKeyValueList(stru interface{})  ([]string ,[]string ){
	newMap := StructToEndMap(stru)

	var list []string

	var value []string


	for k:=range newMap {

		list = append(list,k)
	}
	for z:=range list {
		value=append(value,InterfaceStrval(newMap[list[z]]))
	}
	return list,value

}
func GetStructSqlKeyValueSqlList(stru interface{})  ([]string ,[]string ){
	newMap := StructToEndMap(stru)

	var list []string

	var value []string


	for k:=range newMap {

		list = append(list,k)
	}
	for z:=range list {
		value=append(value,InterfaceStrSqlval(newMap[list[z]]))
	}
	return list,value

}
func GetStructValueList(stru interface{})  []string {
	newMap := StructToEndMap(stru)

	var list []string

	for k:=range newMap {

		list = append(list,InterfaceStrval(newMap[k]))
	}
	return list

}
func GetMapToStructBuff(mapS map[string]interface{},stru interface{}) []byte{

	mapRf := reflect.TypeOf(stru)
	newMaps :=make(map[string]interface{})

	for k,_:=range mapS {
		f ,isHas :=mapRf.FieldByName(k)

		if !isHas {
			bigM :=strings.ToUpper(k[:1])+k[1:]
			f ,isHasTag :=mapRf.FieldByName(bigM)

			if isHasTag {
				switch f.Type.Kind() {
				case reflect.Float64:
					newMaps[k] ,_= strconv.ParseFloat(mapS[k].(string),64)
					break;
				case reflect.Int64:
					newMaps[k] ,_= strconv.Atoi(mapS[k].(string))
					break;
				case reflect.Int:
					newMaps[k] ,_= strconv.Atoi(mapS[k].(string))
					break;
				default:
					newMaps[k] =mapS[k]
				}
			}
		}else {
			switch f.Type.Kind() {
			case reflect.Float64:
				newMaps[k] ,_= strconv.ParseFloat(mapS[k].(string),64)
				break;
			case reflect.Int64:
				newMaps[k] ,_= strconv.Atoi(mapS[k].(string))
				break;
			case reflect.Int:
				newMaps[k] ,_= strconv.Atoi(mapS[k].(string))
				break;
			default:
				newMaps[k] =mapS[k]
			}

		}
	}
	buff ,_:=json.Marshal(newMaps)
	return buff
}



func StructToEndMap(stru interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(stru)
	json.Unmarshal(j, &m)
	return m
}
func GetStructName(info interface{}) string {
	return reflect.TypeOf(info).Name()
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func InterfaceStrSqlval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key ="'"+ value.(string) +"'"
	case []byte:
		key = "'"+string(value.([]byte))+"'"
	default:
		newValue, _ := json.Marshal(value)
		key = "'"+string(newValue)+"'"
	}

	return key
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func InterfaceStrval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}