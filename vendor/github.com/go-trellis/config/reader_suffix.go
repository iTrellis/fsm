// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"strings"
)

type defSuffixReader struct {
	opts   ReaderOptions
	reader Reader
}

// NewSuffixReader return a suffix reader
// supportted: .json, .xml, .yaml, .yml
func NewSuffixReader(opts ...ReaderOptionFunc) (reader Reader, err error) {
	r := &defSuffixReader{}

	for _, o := range opts {
		o(&r.opts)
	}

	if r.opts.filename == "" {
		return nil, ErrInvalidFilePath
	}

	r.reader, err = fileToReader(r.opts.filename)
	if err != nil {
		return
	}
	return r, nil
}

func (p *defSuffixReader) Read(model interface{}) (err error) {
	return p.reader.Read(model)
}

func (p *defSuffixReader) Dump(v interface{}) ([]byte, error) {
	return p.reader.Dump(v)
}

func (p *defSuffixReader) ParseData(data []byte, model interface{}) error {
	return p.reader.ParseData(data, model)
}

func fileToReader(filename string) (Reader, error) {
	switch {
	case strings.HasSuffix(filename, ".json"):
		return NewJSONReader(ReaderOptionFilename(filename)), nil
	case strings.HasSuffix(filename, ".xml"):
		return NewXMLReader(ReaderOptionFilename(filename)), nil
	case strings.HasSuffix(filename, ".yml"),
		strings.HasSuffix(filename, ".yaml"):
		return NewYAMLReader(ReaderOptionFilename(filename)), nil
	default:
		return nil, ErrUnknownSuffixes
	}
}

func fileToReaderType(name string) ReaderType {
	switch {
	case strings.HasSuffix(name, ".json"):
		return ReaderTypeJSON
	case strings.HasSuffix(name, ".xml"):
		return ReaderTypeXML
	case strings.HasSuffix(name, ".yml"),
		strings.HasSuffix(name, ".yaml"):
		return ReaderTypeYAML
	default:
		return ReaderTypeSuffix
	}
}
