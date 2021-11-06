package main

type nullDepartment struct {
	numberOfProfessors int
}

func (c *nullDepartment) getNumberOfProfessors() int {
	return 0
}

func (c *nullDepartment) getName() string {
	return "nullDepartment"
}
