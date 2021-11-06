package main

type college struct {
	departments []department
}

func (c *college) addDepartment(departmentName string, numOfProfessors int) {
	if departmentName == "computerScience" {
		computerScienceDepartment := &computerScience{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, computerScienceDepartment)
	}
	if departmentName == "mechanical" {
		mechanicalDepartment := &mechanical{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, mechanicalDepartment)
	}
	return
}

func (c *college) getDepartment(departmentName string) department {
	for _, department := range c.departments {
		if department.getName() == departmentName {
			return department
		}
	}
	//Return a null department if the department doesn't exits
	return &nullDepartment{}
}
