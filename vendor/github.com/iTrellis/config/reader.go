// GNU GPL v3 License
// Copyright (c) 2017 github.com:iTrellis

package config

import (
	"github.com/iTrellis/common/files"
)

// ReaderType define reader type
type ReaderType int

var filesRepo = files.New()

const (
	// ReaderTypeSuffix judge by file suffix
	ReaderTypeSuffix ReaderType = iota
	// ReaderTypeJSON json reader type
	ReaderTypeJSON
	// ReaderTypeYAML yaml reader type
	ReaderTypeYAML
	// ReaderTypeXML xml reader type
	ReaderTypeXML
)

// Reader reader repo
type Reader interface {
	// read file into model
	Read(model interface{}) error
	// dump configs' cache
	Dump(model interface{}) ([]byte, error)
	// parse data to model
	ParseData(data []byte, model interface{}) error
}

// ReaderOptionFunc declare reader option function
type ReaderOptionFunc func(*ReaderOptions)

//ReaderOptions reader options
type ReaderOptions struct {
	filename string
}

// ReaderOptionFilename set reader filename
func ReaderOptionFilename(filename string) ReaderOptionFunc {
	return func(opts *ReaderOptions) {
		opts.filename = filename
	}
}

// NewReader return a reader by ReaderType
func NewReader(rt ReaderType, filename string) (Reader, error) {
	switch rt {
	case ReaderTypeJSON:
		return NewJSONReader(ReaderOptionFilename(filename)), nil
	case ReaderTypeXML:
		return NewXMLReader(ReaderOptionFilename(filename)), nil
	case ReaderTypeYAML:
		return NewYAMLReader(ReaderOptionFilename(filename)), nil
	default:
		return nil, ErrNotSupportedReaderType
	}
}

/*
SPACE (\u0020)
NO-BREAK SPACE (\u00A0)
OGHAM SPACE MARK (\u1680)
EN QUAD (\u2000)
EM QUAD (\u2001)
EN SPACE (\u2002)
EM SPACE (\u2003)
THREE-PER-EM SPACE (\u2004)
FOUR-PER-EM SPACE (\u2005)
SIX-PER-EM SPACE (\u2006)
FIGURE SPACE (\u2007)
PUNCTUATION SPACE (\u2008)
THIN SPACE (\u2009)
HAIR SPACE (\u200A)
NARROW NO-BREAK SPACE (\u202F)
MEDIUM MATHEMATICAL SPACE (\u205F)
and IDEOGRAPHIC SPACE (\u3000)
Byte Order Mark (\uFEFF)
*/
func isWhitespace(c byte) bool {
	switch string(c) {
	case " ", "\t", "\n", "\u000B", "\u000C",
		"\u000D", "\u00A0", "\u1680", "\u2000",
		"\u2001", "\u2002", "\u2003", "\u2004",
		"\u2005", "\u2006", "\u2007", "\u2008",
		"\u2009", "\u200A", "\u202F", "\u205F",
		"\u2060", "\u3000", "\uFEFF":
		return true
	}
	return false
}
