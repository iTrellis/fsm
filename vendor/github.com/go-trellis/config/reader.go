// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import "io/ioutil"

// ReaderType define reader type
type ReaderType int

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
	Read(path string, model interface{}) error
	// dump configs' cache
	Dump(model interface{}) ([]byte, error)
}

// NewReader return a reader by ReaderType
func NewReader(rt ReaderType) Reader {
	switch rt {
	case ReaderTypeJSON:
		return NewJSONReader()
	case ReaderTypeXML:
		return NewXMLReader()
	case ReaderTypeYAML:
		return NewYAMLReader()
	default:
		return NewSuffixReader()
	}
}

func readFile(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
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
	str := string(c)

	switch str {
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
