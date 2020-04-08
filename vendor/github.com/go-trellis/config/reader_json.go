// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"bytes"
	"encoding/json"
	"sync"
)

type defJSONReader struct {
	mu sync.Mutex
}

var jsonReader = &defJSONReader{}

// NewJSONReader return a json reader
func NewJSONReader() Reader {
	return jsonReader
}

func (p *defJSONReader) Read(name string, model interface{}) error {
	if len(name) == 0 {
		return ErrInvalidFilePath
	}
	p.mu.Lock()
	defer p.mu.Unlock()

	return ReadJSONFile(name, model)
}

func (*defJSONReader) Dump(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// ReadJSONFile 读取Json文件数据到Models
func ReadJSONFile(name string, model interface{}) error {
	data, err := readFile(name)
	if err != nil {
		return err
	}
	return ParseJSONConfig(data, model)
}

// ParseJSONConfig 解析Json配置
func ParseJSONConfig(data []byte, model interface{}) error {
	var escaped bool // string value flag
	var comments int // 1 line; 2 multi line
	var result []byte

	length := len(data)
	for i, w := 0, 0; i < length; i += w {
		w = 1

		switch comments {
		case 1:
			if data[i] == '\n' {
				comments = 0
				escaped = false
			}
			continue
		case 2:
			if data[i] != '*' || length == i+1 {
				continue
			}
			if data[i+1] != '/' {
				continue
			}
			w = 2
			comments = 0
			escaped = false
			continue
		}
		switch data[i] {
		case '"':
			{
				if escaped {
					escaped = false
				} else {
					escaped = true
				}
				result = append(result, data[i])
			}
		case '/':
			{
				if escaped || length == i+1 {
					result = append(result, data[i])
					break
				}
				switch data[i+1] {
				case '/':
					w = 2
					comments = 1
				case '*':
					w = 2
					comments = 2
				default:
					result = append(result, data[i])
				}
			}
		default:
			if escaped || !isWhitespace(data[i]) {
				result = append(result, data[i])
			}

		}
	}

	decoder := json.NewDecoder(bytes.NewBuffer(result))
	decoder.UseNumber()

	return decoder.Decode(model)
}
