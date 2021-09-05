package main

type cabinBuilder struct {
	house
}

func newCabinBuilder() *cabinBuilder {
	return &cabinBuilder{}
}

func (b *cabinBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *cabinBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *cabinBuilder) setNumFloor() {
	b.floor = 2
}

func (b *cabinBuilder) getHouse() house {
	return b.house
}
