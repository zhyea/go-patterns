package main

import "fmt"

func main() {
	college1 := createCollege1()
	college2 := createCollege2()
	totalProfessors := 0
	departmentArray := []string{"computerScience", "mechanical", "civil", "electronics"}

	for _, departmentName := range departmentArray {
		d := college1.getDepartment(departmentName)
		totalProfessors += d.getNumberOfProfessors()
	}

	fmt.Printf("Total number of professors in college1 is %d\n", totalProfessors)

	//Reset the professor count
	totalProfessors = 0
	for _, departmentName := range departmentArray {
		d := college2.getDepartment(departmentName)
		totalProfessors += d.getNumberOfProfessors()
	}
	fmt.Printf("Total number of professors in college2 is %d\n", totalProfessors)
}

func createCollege1() *college {
	college := &college{}
	college.addDepartment("computerScience", 4)
	college.addDepartment("mechanical", 5)
	return college
}

func createCollege2() *college {
	college := &college{}
	college.addDepartment("computerScience", 2)
	return college
}
