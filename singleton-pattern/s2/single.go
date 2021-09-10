package s2

import (
	"fmt"
	"log"
)

type single struct {
}

var singleInstance *single

func init() {
	fmt.Println("Creating Single Instance Now")
	singleInstance = &single{}
}

func GetInstance() *single {
	if singleInstance == nil {
		log.Fatal("Single Instance is nil")
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}
