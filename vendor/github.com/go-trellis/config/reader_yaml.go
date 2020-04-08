// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"sync"

	"gopkg.in/yaml.v2"
)

type defYamlReader struct {
	mu sync.Mutex
}

var yamlReader = &defYamlReader{}

// NewYAMLReader return a yaml reader
func NewYAMLReader() Reader {
	return yamlReader
}

func (p *defYamlReader) Read(name string, model interface{}) error {
	if len(name) == 0 {
		return ErrInvalidFilePath
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	data, err := ReadYAMLFile(name)
	if err != nil {
		return err
	}
	return ParseYAMLConfig(data, model)
}

func (*defYamlReader) Dump(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// ReadYAMLFile 读取yaml文件的配置信息
func ReadYAMLFile(name string) ([]byte, error) {
	return readFile(name)
}

// ParseYAMLConfig 解析yaml的配置信息
func ParseYAMLConfig(data []byte, model interface{}) error {
	return yaml.Unmarshal(data, model)
}
