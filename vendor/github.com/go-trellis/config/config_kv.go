// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"reflect"
	"strings"

	"github.com/go-trellis/common/formats"
)

func (p *AdapterConfig) copyJSONDollarSymbol(key string, maps *map[string]interface{}) error {
	tokens := []string{}
	if key != "" {
		tokens = append(tokens, key)
	}
	for k, v := range *maps {
		if v == nil {
			return nil
		}
		keys := append(tokens, k)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			{
				vm, ok := v.(map[string]interface{})
				if !ok {
					continue
				}
				err := p.copyJSONDollarSymbol(strings.Join(keys, "."), &vm)
				if err != nil {
					return err
				}
			}
		case reflect.String:
			{
				s, ok := v.(string)
				if !ok {
					continue
				}
				_, matched := formats.FindStringSubmatchMap(s, includeReg)
				if !matched {
					continue
				}

				vm, err := p.getKeyValue(s[2 : len(s)-1])
				if err != nil {
					return err
				}
				err = p.setKeyValue(strings.Join(keys, "."), vm)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (p *AdapterConfig) copyYAMLDollarSymbol(configs *map[string]interface{}) error {

	for k, v := range *configs {
		if v == nil {
			return nil
		}
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			{
				vm, ok := v.(map[interface{}]interface{})
				if !ok {
					continue
				}
				if err := p.copyMap(k, &vm); err != nil {
					return err
				}
			}
		case reflect.String:
			{
				s, ok := v.(string)
				if !ok {
					continue
				}
				if _, matched := formats.FindStringSubmatchMap(s, includeReg); !matched {
					continue
				}
				vm, err := p.getKeyValue(s[2 : len(s)-1])
				if err != nil {
					return err
				}
				err = p.setKeyValue(k, vm)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (p *AdapterConfig) copyMap(key string, maps *map[interface{}]interface{}) error {
	tokens := []string{}
	if key != "" {
		tokens = append(tokens, key)
	}

	for k, v := range *maps {
		if v == nil {
			return nil
		}
		keys := append(tokens, k.(string))
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			{
				vm, ok := v.(map[interface{}]interface{})
				if !ok {
					continue
				}
				err := p.copyMap(strings.Join(keys, "."), &vm)
				if err != nil {
					return err
				}
				err = p.setKeyValue(strings.Join(keys, "."), vm)
				if err != nil {
					return err
				}
			}
		case reflect.String:
			{
				s, ok := v.(string)
				if !ok {
					continue
				}
				if _, matched := formats.FindStringSubmatchMap(s, includeReg); !matched {
					continue
				}
				vm, e := p.getKeyValue(s[2 : len(s)-1])
				if e != nil {
					continue
				}
				err := p.setKeyValue(strings.Join(keys, "."), vm)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (p *AdapterConfig) getKeyValue(key string) (vm interface{}, err error) {

	tokens := strings.Split(key, ".")
	vm = p.configs[tokens[0]]
	for i, t := range tokens {
		if i == 0 {
			continue
		}

		switch v := vm.(type) {
		case map[string]interface{}:
			vm = v[t]
		case map[interface{}]interface{}:
			vm = v[t]
		default:
			return nil, ErrNotMap
		}

	}
	if vm == nil {
		err = ErrValueNil
	}
	return
}

// setKeyValue set key value into *configs
func (p *AdapterConfig) setKeyValue(key string, value interface{}) (err error) {
	tokens := strings.Split(key, ".")
	for i := len(tokens) - 1; i >= 0; i-- {
		if i == 0 {
			p.configs[tokens[0]] = value
			return
		}
		v, _ := p.getKeyValue(strings.Join(tokens[:i], "."))
		switch vm := v.(type) {
		case map[string]interface{}:
			vm[tokens[i]] = value
			value = vm
		case map[interface{}]interface{}:
			vm[tokens[i]] = value
			value = vm
		default:
			value = map[string]interface{}{tokens[i]: value}
		}
	}
	return
}
