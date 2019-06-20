// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package fsm

import (
	"github.com/gogap/config"
)

// NewTransactionFromConfig new transactions from config file
func NewTransactionFromConfig(filepath string) error {
	return NewTransactions(config.NewConfig(config.ConfigFile(filepath)))
}

// NewTransactions new transactions
func NewTransactions(cfg config.Configuration) (err error) {
	f := New()
	fsmConfig := cfg.GetConfig("fsm")
	for _, namespace := range fsmConfig.Keys() {
		nsConfig := fsmConfig.GetConfig(namespace)
		for _, key := range nsConfig.Keys() {
			obj := nsConfig.GetConfig(key)
			f.Add(&Transaction{
				Namespace:     namespace,
				CurrentStatus: obj.GetString("current"),
				Event:         obj.GetString("event"),
				TargetStatus:  obj.GetString("target"),
			})
		}
	}
	return
}
