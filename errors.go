// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package fsm

import (
	"github.com/go-rut/errors"
)

const (
	namespace = "Trellis::FSM"
)

var (
	ErrInvalidTransaction = errors.TN(namespace, 1000, "invalid transaction")
	ErrTargetStatusEmpty  = errors.TN(namespace, 1001, "empty target status")
)
