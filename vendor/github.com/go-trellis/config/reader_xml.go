// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"encoding/xml"
)

type defXMLReader struct {
	opts ReaderOptions
}

// NewXMLReader return xml config reader
func NewXMLReader(opts ...ReaderOptionFunc) Reader {
	r := &defXMLReader{}
	for _, o := range opts {
		o(&r.opts)
	}
	return r
}

func (p *defXMLReader) Read(model interface{}) error {
	data, err := ReadXMLFile(p.opts.filename)
	if err != nil {
		return err
	}
	return ParseXMLConfig(data, model)
}

func (*defXMLReader) Dump(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (*defXMLReader) ParseData(data []byte, model interface{}) error {
	return ParseXMLConfig(data, model)
}

// ReadXMLFile 读取yaml文件的配置信息
func ReadXMLFile(name string) ([]byte, error) {
	data, _, err := filesRepo.Read(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ParseXMLConfig 解析yaml的配置信息
func ParseXMLConfig(data []byte, model interface{}) error {
	return xml.Unmarshal(data, model)
}
