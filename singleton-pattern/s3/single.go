package s3

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating Single Instance Now")
				singleInstance = &single{}
			})
		fmt.Println("Single Instance already created-1")
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}
