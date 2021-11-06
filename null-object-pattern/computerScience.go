package main

type computerScience struct {
	numberOfProfessors int
}

func (c *computerScience) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *computerScience) getName() string {
	return "computerScience"
}
