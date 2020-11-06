// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"bytes"
	"encoding/json"
)

type defJSONReader struct {
	opts ReaderOptions
}

// NewJSONReader return a json reader
func NewJSONReader(opts ...ReaderOptionFunc) Reader {
	r := &defJSONReader{}
	for _, o := range opts {
		o(&r.opts)
	}
	return r
}

func (p *defJSONReader) Read(model interface{}) error {
	return ReadJSONFile(p.opts.filename, model)
}

func (*defJSONReader) Dump(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (*defJSONReader) ParseData(data []byte, model interface{}) error {
	return ParseJSONConfig(data, model)
}

// ReadJSONFile 读取Json文件数据到Models
func ReadJSONFile(filepath string, model interface{}) error {
	data, _, err := filesRepo.Read(filepath)
	if err != nil {
		return err
	}
	return ParseJSONConfig(data, model)
}

// ParseJSONConfig 解析Json配置
func ParseJSONConfig(data []byte, model interface{}) error {
	var escaped bool // string value flag, " appear times, odd is false, even is true
	var comments int // 0 nothing; 1 line; 2 multi line
	var result []byte

	length := len(data)
	for i, w := 0, 0; i < length; i += w {
		w = 1

		switch comments {
		case 1:
			if data[i] == '\n' {
				comments, escaped = 0, false
			}
			continue
		case 2:
			// */
			if data[i] == '*' && length != i+1 && data[i+1] == '/' {
				w = 2
				comments, escaped = 0, false
			}
			continue
		}

		switch data[i] {
		case '"':
			escaped = !escaped
			result = append(result, data[i])
		case '/':
			if escaped || length == i+1 {
				result = append(result, data[i])
				continue
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
