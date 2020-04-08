// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"errors"
)

// Errors
var (
	ErrNotMap                 = errors.New("interface not map")
	ErrValueNil               = errors.New("value is nil")
	ErrInvalidKey             = errors.New("invalid key")
	ErrInvalidFilePath        = errors.New("invalid file path")
	ErrUnknownSuffixes        = errors.New("unknown file with suffix")
	ErrNotSupportedReaderType = errors.New("not supported reader type")
)
