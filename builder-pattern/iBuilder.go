package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "cabin" {
		return &cabinBuilder{}
	}
	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}
