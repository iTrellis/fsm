// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"gopkg.in/yaml.v2"
)

type defYamlReader struct {
	opts ReaderOptions
}

// NewYAMLReader return a yaml reader
func NewYAMLReader(opts ...ReaderOptionFunc) Reader {
	r := &defYamlReader{}
	for _, o := range opts {
		o(&r.opts)
	}
	return r
}

func (p *defYamlReader) Read(model interface{}) error {
	data, err := ReadYAMLFile(p.opts.filename)
	if err != nil {
		return err
	}
	return ParseYAMLConfig(data, model)
}

func (*defYamlReader) Dump(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (*defYamlReader) ParseData(data []byte, model interface{}) error {
	return ParseYAMLConfig(data, model)
}

// ReadYAMLFile 读取yaml文件的配置信息
func ReadYAMLFile(name string) ([]byte, error) {
	data, _, err := filesRepo.Read(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ParseYAMLConfig 解析yaml的配置信息
func ParseYAMLConfig(data []byte, model interface{}) error {
	return yaml.Unmarshal(data, model)
}
