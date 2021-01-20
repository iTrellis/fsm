/*
Copyright © 2017 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

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
