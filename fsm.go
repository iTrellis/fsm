// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package fsm

import (
	"sync"
)

type fsm struct {
	Transations map[string]map[string]*Transaction

	sync.RWMutex
}

var defaultFSM *fsm

// New get default fsm
func New() FSMRepo {
	if defaultFSM == nil {
		defaultFSM = &fsm{
			Transations: make(map[string]map[string]*Transaction, 0),
		}
	}
	return defaultFSM
}

// Add add a transaction
func (p *fsm) Add(t *Transaction) {
	if e := t.valid(); e != nil {
		return
	}

	p.Lock()
	defer p.Unlock()
	p.add(t)
}

func (p *fsm) add(t *Transaction) {

	spaceTrans := p.Transations[t.Namespace]

	if spaceTrans == nil {
		spaceTrans = make(map[string]*Transaction)
	}

	spaceTrans[p.genKey(t.CurrentStatus, t.Event)] = t
	p.Transations[t.Namespace] = spaceTrans
}

// GetTargetTranstion get trans by current information
func (p *fsm) GetTargetTranstion(namespace, curStatus, event string) *Transaction {
	p.RLock()
	defer p.RUnlock()
	return p.getTransaction(namespace, curStatus, event)
}

// Remove remove all transactions
func (p *fsm) Remove() {
	p.Lock()
	defer p.Unlock()
	p.remove()
}

func (p *fsm) remove() {
	p.Transations = make(map[string]map[string]*Transaction, 0)
}

// RemoveNamespace remove namespace's transactions
func (p *fsm) RemoveNamespace(namespace string) {
	if namespace == "" {
		return
	}
	p.Lock()
	defer p.Unlock()
	p.removeNamespace(namespace)
}

func (p *fsm) removeNamespace(namespace string) {
	delete(p.Transations, namespace)
}

// RemoveByTransaction remove a transaction by current information
func (p *fsm) RemoveByTransaction(t *Transaction) {
	if e := t.validCurrent(); e != nil {
		return
	}
	p.Lock()
	defer p.Unlock()
	p.removeByTransaction(t)
}

func (p *fsm) removeByTransaction(t *Transaction) {
	spaceTrans := p.Transations[t.Namespace]
	delete(spaceTrans, p.genKey(t.CurrentStatus, t.Event))
	p.Transations[t.Namespace] = spaceTrans
}

func (p *fsm) getTransaction(namespace, curStatus, event string) *Transaction {
	spaceTrans := p.Transations[namespace]
	if spaceTrans == nil {
		return nil
	}
	return spaceTrans[p.genKey(curStatus, event)]
}

func (p *fsm) genKey(curStatus, event string) string {
	return curStatus + "::" + event
}
