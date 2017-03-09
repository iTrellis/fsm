// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package fsm

import (
	"errors"
)

// errors
var (
	ErrInvalidTransaction = errors.New("invalid transaction")
	ErrTargetStatusEmpty  = errors.New("empty target status")
)
