package utils

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

/*
	由map拼接k=v字符串
*/
func BuildMapToSortString(m map[string]interface{}) (str string) {
	//key=value&key1=value2...
	EachMap(m, func(key string, value interface{}) {
		str += fmt.Sprintf("%s=%v&", key, value)
	})
	str = str[:(len(str)-1)]
	return str
}

/*
	map排序
*/
func EachMap(eachMap interface{}, eachFunc interface{})  {
	eachMapValue := reflect.ValueOf(eachMap)
	eachFuncValue := reflect.ValueOf(eachFunc)
	eachMapType := eachMapValue.Type()
	eachFuncType := eachFuncValue.Type()
	if eachMapValue.Kind() != reflect.Map {
		panic(errors.New("ksort.EachMap failed. parameter \"eachMap\" type must is map[...]...{}"))
	}
	if eachFuncValue.Kind() != reflect.Func {
		panic(errors.New("ksort.EachMap failed. parameter \"eachFunc\" type must is func(key ..., value ...)"))
	}
	if eachFuncType.NumIn() != 2 {
		panic(errors.New("ksort.EachMap failed. \"eachFunc\" input parameter count must is 2"))
	}
	if eachFuncType.In(0).Kind() != eachMapType.Key().Kind() {
		panic(errors.New("ksort.EachMap failed. \"eachFunc\" input parameter 1 type not equal of \"eachMap\" key"))
	}
	if eachFuncType.In(1).Kind() != eachMapType.Elem().Kind() {
		panic(errors.New("ksort.EachMap failed. \"eachFunc\" input parameter 2 type not equal of \"eachMap\" value"))
	}

	// 对key进行排序
	// 获取排序后map的key和value，作为参数调用eachFunc即可
	switch eachMapType.Key().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		keys := make([]int, 0)
		keysMap := map[int]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, int(value.Int()))
			keysMap[int(value.Int())] = value
		}
		sort.Ints(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	case reflect.Float64, reflect.Float32:
		keys := make([]float64, 0)
		keysMap := map[float64]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, float64(value.Float()))
			keysMap[float64(value.Float())] = value
		}
		sort.Float64s(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	case reflect.String:
		keys := make([]string, 0)
		keysMap := map[string]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, value.String())
			keysMap[value.String()] = value
		}
		sort.Strings(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	default:
		panic(errors.New("\"eachMap\" key type must is int or float or string"))
	}
}

