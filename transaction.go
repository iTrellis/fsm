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

// Transaction information for current to target status in namespace
type Transaction struct {
	Namespace     string `json:"namespace"`
	CurrentStatus string `json:"current"`
	Event         string `json:"event"`
	TargetStatus  string `json:"target"`
}

func (p *Transaction) valid() error {

	if e := p.validCurrent(); e != nil {
		return e
	}

	if p.TargetStatus == "" {
		return ErrTargetStatusEmpty
	}

	return nil
}

func (p *Transaction) validCurrent() error {

	if p == nil {
		return ErrInvalidTransaction
	}

	if p.Namespace == "" ||
		p.Event == "" ||
		p.CurrentStatus == "" {
		return ErrInvalidTransaction
	}
	return nil
}
