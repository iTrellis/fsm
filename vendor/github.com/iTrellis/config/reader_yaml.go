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
