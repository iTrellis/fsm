// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"encoding/xml"
	"sync"
)

type defXMLReader struct {
	mu sync.Mutex
}

var xmlReader = &defXMLReader{}

// NewXMLReader return xml config reader
func NewXMLReader() Reader {
	return xmlReader
}

func (p *defXMLReader) Read(name string, model interface{}) error {
	if name == "" {
		return ErrInvalidFilePath
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	data, err := ReadXMLFile(name)
	if err != nil {
		return err
	}
	return ParseXMLConfig(data, model)
}

func (*defXMLReader) Dump(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

// ReadXMLFile 读取yaml文件的配置信息
func ReadXMLFile(name string) ([]byte, error) {
	return readFile(name)
}

// ParseXMLConfig 解析yaml的配置信息
func ParseXMLConfig(data []byte, model interface{}) error {
	return xml.Unmarshal(data, model)
}
