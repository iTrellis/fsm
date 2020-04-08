// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package formats

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Strings array string
type Strings []string

func (x Strings) Len() int { return len(x) }

func (x Strings) Less(i, j int) bool { return x[i] < x[j] }

func (x Strings) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// IntToString parse int to string
func IntToString(value interface{}) (string, error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int8:
		return fmt.Sprintf("%d", value.(int8)), nil
	case reflect.Int16:
		return fmt.Sprintf("%d", value.(int16)), nil
	case reflect.Int:
		return fmt.Sprintf("%d", value.(int)), nil
	case reflect.Int32:
		return fmt.Sprintf("%d", value.(int32)), nil
	case reflect.Int64:
		return strconv.FormatInt(value.(int64), 10), nil
	case reflect.Uint:
		return fmt.Sprintf("%d", value.(uint)), nil
	case reflect.Uint8:
		return fmt.Sprintf("%d", value.(uint8)), nil
	case reflect.Uint16:
		return fmt.Sprintf("%d", value.(uint16)), nil
	case reflect.Uint32:
		return fmt.Sprintf("%d", value.(uint32)), nil
	case reflect.Uint64:
		return fmt.Sprintf("%d", value.(uint64)), nil
	default:
		return "", fmt.Errorf("type is valid: %s", reflect.TypeOf(value).String())
	}
}

// HideString hide some words
// origin: the string to be hidden
// start:  from 1 to len(origin), replace * from beginning
// length: replace origin length from the beginning
func HideString(origin string, start, length int) string {

	if len(origin) == 0 || length <= 0 || start <= 0 {
		return origin
	}

	start--

	rs := []rune(origin)
	lenRs := len(rs)

	if lenRs >= start {
		if lenRs > length+start {
			return string(rs[0:start]) + strings.Repeat("*", length) + string(rs[length+start:])
		}
		return string(rs[0:start]) + strings.Repeat("*", lenRs-start)
	}
	return origin
}
