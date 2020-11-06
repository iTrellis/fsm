// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"math/big"
	"time"
)

// OptionFunc 处理函数
type OptionFunc func(*AdapterConfig)

// OptionFile 解析配置文件Option函数
func OptionFile(filename string) OptionFunc {
	return func(c *AdapterConfig) {
		c.ConfigFile = filename
	}
}

// OptionString 字符串解析配置Option函数
func OptionString(rt ReaderType, cStr string) OptionFunc {
	return func(c *AdapterConfig) {
		c.readerType = rt
		c.ConfigString = cStr
	}
}

// OptionStruct 结构体解析配置Option函数
func OptionStruct(rt ReaderType, st interface{}) OptionFunc {
	return func(c *AdapterConfig) {
		c.readerType = rt
		c.ConfigStruct = st
	}
}

// Config manager data functions
type Config interface {
	// get a object
	GetInterface(key string, defValue ...interface{}) (res interface{})
	// get a string
	GetString(key string, defValue ...string) (res string)
	// get a bool
	GetBoolean(key string, defValue ...bool) (b bool)
	// get a int
	GetInt(key string, defValue ...int) (res int)
	// get a float
	GetFloat(key string, defValue ...float64) (res float64)
	// get list of objects
	GetList(key string) (res []interface{})
	// get list of strings
	GetStringList(key string) []string
	// get list of bools
	GetBooleanList(key string) []bool
	// get list of ints
	GetIntList(key string) []int
	// get list of float64s
	GetFloatList(key string) []float64
	// get time duration by (int)(uint), exp: 1s, 1day
	GetTimeDuration(key string, defValue ...time.Duration) time.Duration
	// get byte size by (int)(uint), exp: 1k, 1m
	GetByteSize(key string) *big.Int
	// get map value
	GetMap(key string) Options
	// get key's config
	GetConfig(key string) Config
	// ToObject unmarshal values to object
	ToObject(key string, model interface{}) error
	// get key's values if values can be Config, or panic
	GetValuesConfig(key string) Config
	// set key's value into config
	SetKeyValue(key string, value interface{}) (err error)
	// get all config
	Dump() (bs []byte, err error)
	// get all keys
	GetKeys() []string
	// deep copy configs
	Copy() Config
}

// NewConfig return Config by file's path, judge path's suffix, supported .json, .yml, .yaml
func NewConfig(name string) (Config, error) {
	if len(name) == 0 {
		return nil, ErrInvalidFilePath
	}
	return NewConfigOptions(OptionFile(name))
}

// NewConfigOptions 从操作函数解析Config
func NewConfigOptions(opts ...OptionFunc) (Config, error) {
	c := &AdapterConfig{}
	err := c.init(opts...)
	if err != nil {
		return nil, err
	}

	return c, nil
}
