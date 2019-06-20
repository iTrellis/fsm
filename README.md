# fsm
Finite-state machine in go

* [![Build Status](https://travis-ci.org/go-trellis/fsm.png)](https://travis-ci.org/go-trellis/fsm)
* [![GoDoc](http://godoc.org/github.com/go-trellis/fsm?status.svg)](http://godoc.org/github.com/go-trellis/fsm)

## Introduction

* [点击进入中文相关说明](http://zh.wikipedia.org/wiki/%E6%9C%89%E9%99%90%E7%8A%B6%E6%80%81%E6%9C%BA)
* [Click to article in English](http://en.wikipedia.org/wiki/Finite-state_machine)

## Installation

```go
go get -u github.com/go-trellis/fsm
```

## Usage

### fsm repo

```go
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
```

### new and input a namespace's transaction

```go
	f := fsm.New()

	f.Add(&fsm.Transaction{
			Namespace:     "namespace",
			CurrentStatus: "status1",
			Event:         "event1",
			TargetStatus:  "status2",
		})
	fmt.Println(f.GetTargetTranstion("namespace", "status1", "event1"))

	f.Remove()

	fmt.Println(f.GetTargetTranstion("namespace", "status1", "event1"))
```

## Config

* [sample.conf](sample.conf)