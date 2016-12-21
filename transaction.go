// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package fsm

type Transaction struct {
	Namespace     string
	CurrentStatus string
	Event         string
	TargetStatus  string
}

func (p *Transaction) valid() error {

	if e := p.validCurrent(); e != nil {
		return e
	}

	if p.TargetStatus == "" {
		return ErrTargetStatusEmpty.New()
	}

	return nil
}

func (p *Transaction) validCurrent() error {

	if p == nil {
		return ErrInvalidTransaction.New()
	}

	if p.Namespace == "" ||
		p.Event == "" ||
		p.CurrentStatus == "" {
		return ErrInvalidTransaction.New()
	}
	return nil
}
