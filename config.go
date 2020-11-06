/*
Copyright © 2016 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

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
