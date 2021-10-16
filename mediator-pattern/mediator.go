package main

type mediator interface {
	canLand(train) bool

	notifyFree()
}
