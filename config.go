// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package fsm

import (
	"github.com/go-trellis/config"
)

// NewTransactionFromConfig new transactions from config file
func NewTransactionFromConfig(filepath string) error {
	cfg, err := config.NewConfigOptions(config.OptionFile(filepath))
	if err != nil {
		return err
	}
	return NewTransactions(cfg)
}

// NewTransactions new transactions
func NewTransactions(cfg config.Config) (err error) {
	f := New()
	fsmConfig := cfg.GetValuesConfig("fsm")
	for _, namespace := range fsmConfig.GetKeys() {
		nsConfig := fsmConfig.GetValuesConfig(namespace)
		for _, key := range nsConfig.GetKeys() {
			obj := nsConfig.GetValuesConfig(key)
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
