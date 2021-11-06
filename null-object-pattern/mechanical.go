package main

type mechanical struct {
	numberOfProfessors int
}

func (c *mechanical) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *mechanical) getName() string {
	return "mechanical"
}
