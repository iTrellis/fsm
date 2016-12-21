// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package fsm

// FSMRepo the functions of fsm interface
type FSMRepo interface {
	// add a transction into cache
	Add(*Transaction)
	// remove all transactions
	Remove()
	// remove namespace's transactions
	RemoveNamespace(namespace string)
	// remove a transaction by information
	RemoveByTransaction(*Transaction)
	// get target transaction by current information
	GetTargetTranstion(namespace, curStatus, event string) *Transaction
}
