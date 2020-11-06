// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

// Options initial params
type Options map[string]interface{}

// ToConfig Options to config, default YAML reader
func (p *Options) ToConfig(rts ...ReaderType) Config {
	rt := ReaderTypeYAML
	if len(rts) > 0 {
		rt = rts[0]
	}
	c := &AdapterConfig{readerType: rt, configs: *p}
	switch c.readerType {
	case ReaderTypeJSON:
		c.reader = NewJSONReader()
	case ReaderTypeYAML:
		c.reader = NewYAMLReader()
	default:
		return nil
	}
	return c
}
