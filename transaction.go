// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

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
