package main

type train interface {
	requestArrival()

	departure()

	permitArrival()
}
