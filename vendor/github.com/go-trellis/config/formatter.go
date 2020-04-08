// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"fmt"
	"time"

	"github.com/go-trellis/common/formats"
)

// MapValueGetter get a value from map
type MapValueGetter interface {
	GetMapKeyValue(ms interface{}, key string) (interface{}, error)
	GetMapKeyValueList(ms interface{}, key string) ([]interface{}, error)

	GetMapKeyValueString(ms interface{}, key string) (string, error)
	GetMapKeyValueStringList(ms interface{}, key string) ([]string, error)

	GetMapKeyValueInt(ms interface{}, key string) (int, error)
	GetMapKeyValueIntList(ms interface{}, key string) ([]int, error)

	GetMapKeyValueInt64(ms interface{}, key string) (int64, error)
	GetMapKeyValueInt64List(ms interface{}, key string) ([]int64, error)

	GetMapKeyValueBool(ms interface{}, key string) (bool, error)
	GetMapKeyValueBoolList(ms interface{}, key string) ([]bool, error)
	GetMapKeyValueTimeDuration(ms interface{}, key string) (time.Duration, error)

	GenMapConfig(rt ReaderType, ms map[string]interface{}) Config
}

var defaultGetter = MapGetter()

// MapGetter get map value getter
func MapGetter() MapValueGetter {
	return (*getter)(nil)
}

type getter struct{}

// GetMapKeyValue get value from map[interface{} | string] interface{}
func (*getter) GetMapKeyValue(ms interface{}, key string) (interface{}, error) {
	m, ok := ms.(map[interface{}]interface{})
	if ok {
		return m[key], nil
	}

	s, ok := ms.(map[string]interface{})
	if ok {
		return s[key], nil
	}

	o, ok := ms.(Options)
	if ok {
		return o[key], nil
	}

	return nil, fmt.Errorf("config invalid: %v", ms)
}

func (p *getter) GetMapKeyValueString(ms interface{}, key string) (string, error) {
	v, e := p.GetMapKeyValue(ms, key)
	if e != nil {
		return "", e
	}
	if v == nil {
		return "", nil
	}
	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("value is not string")
	}

	return s, nil
}

func (p *getter) GetMapKeyValueInt(ms interface{}, key string) (int, error) {
	v, e := p.GetMapKeyValue(ms, key)
	if e != nil {
		return 0, e
	}
	if v == nil {
		return 0, nil
	}
	return formats.ToInt(v)
}

func (p *getter) GetMapKeyValueInt64(ms interface{}, key string) (int64, error) {
	v, e := p.GetMapKeyValue(ms, key)
	if e != nil {
		return 0, e
	}
	if v == nil {
		return 0, nil
	}
	return formats.ToInt64(v)
}

func (p *getter) GetMapKeyValueBool(ms interface{}, key string) (bool, error) {
	v, e := p.GetMapKeyValue(ms, key)
	if e != nil {
		return false, e
	}
	if v == nil {
		return false, nil
	}
	b, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf("value is not bool")
	}
	return b, nil
}

func (p *getter) GetMapKeyValueTimeDuration(ms interface{}, key string) (time.Duration, error) {
	s, e := p.GetMapKeyValueString(ms, key)
	if e != nil {
		return 0, e
	}

	return formats.ParseStringTime(s, 0), nil
}

func (*getter) GenMapConfig(rt ReaderType, ms map[string]interface{}) Config {

	c := &AdapterConfig{readerType: rt, reader: nil, configs: ms}
	switch rt {
	case ReaderTypeJSON:
		c.reader = NewJSONReader()
	case ReaderTypeYAML:
		c.reader = NewYAMLReader()
	default:
		return nil
	}
	return c
}
