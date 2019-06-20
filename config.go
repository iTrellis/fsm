package fsm

import (
	"github.com/go-akka/configuration"
)

// NewTransactionFromConfig var loggers map[string]LoggerWriter
func NewTransactionFromConfig(filepath string) error {
	return NewTransactions(configuration.LoadConfig(filepath))
}

// NewTransactions var loggers map[string]LoggerWriter
func NewTransactions(cfg *configuration.Config) (err error) {
	f := New()
	for _, namespace := range cfg.GetNode("fsm").GetObject().GetKeys() {
		namespaceConfig := cfg.GetValue("fsm." + namespace)
		for _, v := range namespaceConfig.GetArray() {
			obj := v.GetObject()
			f.Add(&Transaction{
				Namespace:     namespace,
				CurrentStatus: obj.GetKey("current").GetString(),
				Event:         obj.GetKey("event").GetString(),
				TargetStatus:  obj.GetKey("target").GetString(),
			})
		}
	}
	return
}
