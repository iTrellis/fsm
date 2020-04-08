// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"fmt"
	"reflect"

	"github.com/go-trellis/common/formats"
)

func (p *getter) GetMapKeyValueList(ms interface{}, key string) ([]interface{}, error) {
	v, e := p.GetMapKeyValue(ms, key)
	if e != nil {
		return nil, e
	}

	if v == nil {
		return nil, nil
	}

	vs, ok := v.([]interface{})
	if !ok {
		return nil, fmt.Errorf("value is not slice")
	}

	return vs, nil
}

func (p *getter) GetMapKeyValueStringList(ms interface{}, key string) ([]string, error) {
	vs, e := p.GetMapKeyValueList(ms, key)
	if e != nil {
		return nil, e
	} else if vs == nil {
		return nil, nil
	}

	var ss []string
	for _, v := range vs {
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("value(%s) is not slice of string: %s", key, reflect.TypeOf(v).String())
		}
		ss = append(ss, s)
	}
	return ss, nil
}

func (p *getter) GetMapKeyValueIntList(ms interface{}, key string) ([]int, error) {
	vs, e := p.GetMapKeyValueList(ms, key)
	if e != nil {
		return nil, e
	} else if vs == nil {
		return nil, nil
	}

	var is []int
	for _, v := range vs {
		i, e := formats.ToInt(v)
		if e != nil {
			return nil, e
		}
		is = append(is, i)
	}
	return is, nil
}

func (p *getter) GetMapKeyValueInt64List(ms interface{}, key string) ([]int64, error) {
	vs, e := p.GetMapKeyValueList(ms, key)
	if e != nil {
		return nil, e
	} else if vs == nil {
		return nil, nil
	}

	var is []int64
	for _, v := range vs {
		i, e := formats.ToInt64(v)
		if e != nil {
			return nil, e
		}
		is = append(is, i)
	}
	return is, nil
}

func (p *getter) GetMapKeyValueBoolList(ms interface{}, key string) ([]bool, error) {
	vs, e := p.GetMapKeyValueList(ms, key)
	if e != nil {
		return nil, e
	} else if vs == nil {
		return nil, nil
	}

	var bs []bool
	for _, v := range vs {
		b, ok := v.(bool)
		if !ok {
			return nil, fmt.Errorf("value(%s) is not slice of bool: %s", key, reflect.TypeOf(v).String())
		}
		bs = append(bs, b)
	}
	return bs, nil
}
