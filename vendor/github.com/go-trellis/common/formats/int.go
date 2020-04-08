// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package formats

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// Uint32s array uint32
type Uint32s []uint32

// Ints array int
type Ints []int

// Int64s array int64
type Int64s []int64

// Uints array uint
type Uints []uint

func (x Uint32s) Len() int { return len(x) }

func (x Uint32s) Less(i, j int) bool { return x[i] < x[j] }

func (x Uint32s) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x Uints) Len() int { return len(x) }

func (x Uints) Less(i, j int) bool { return x[i] < x[j] }

func (x Uints) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x Ints) Len() int { return len(x) }

func (x Ints) Less(i, j int) bool { return x[i] < x[j] }

func (x Ints) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x Int64s) Len() int { return len(x) }

func (x Int64s) Less(i, j int) bool { return x[i] < x[j] }

func (x Int64s) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// ToInt64 parse value to int64
func ToInt64(value interface{}) (int64, error) {
	var val string
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		val = fmt.Sprintf("%d", value)
	case reflect.String:
		switch reflect.TypeOf(value).String() {
		case "json.Number":
			return value.(json.Number).Int64()
		default:
			val = value.(string)
		}
	default:
		return 0, fmt.Errorf("type is valid: %s", reflect.TypeOf(value).String())
	}

	return strconv.ParseInt(val, 10, 64)
}

// ToInt parse value to int
func ToInt(value interface{}) (int, error) {
	var val string
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		val = fmt.Sprintf("%d", value)
	case reflect.String:
		switch reflect.TypeOf(value).String() {
		case "json.Number":
			val = value.(json.Number).String()
		default:
			val = value.(string)
		}
	default:
		return 0, fmt.Errorf("type is valid: %s", reflect.TypeOf(value).String())
	}

	return strconv.Atoi(val)
}
