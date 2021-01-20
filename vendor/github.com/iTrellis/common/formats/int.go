/*
Copyright © 2020 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

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
