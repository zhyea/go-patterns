package main

import (
	"fmt"
)

func main() {
	cabinBuilder := getBuilder("cabin")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(cabinBuilder)
	cabinHouse := director.buildHouse()

	fmt.Printf("Cabin House Door Type: %s\n", cabinHouse.doorType)
	fmt.Printf("Cabin House Window Type: %s\n", cabinHouse.windowType)
	fmt.Printf("Cabin House Num Floor: %d\n", cabinHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
