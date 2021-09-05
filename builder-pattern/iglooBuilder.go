package main

type iglooBuilder struct {
	house
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Ice Window"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Ice Door"
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() house {
	return b.house
}
