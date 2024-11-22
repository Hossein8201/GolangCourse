package __Closures

import (
	"fmt"
	"time"
)

func Example() {
	myString := "Ali"
	myFunc := func() {
		fmt.Println(myString)
	}
	myString = "Reza"
	myFunc()
	//
	mySlice := []string{"Ali", "Reza", "Amir", "Hossein"}
	for index, name := range mySlice {
		func() {
			fmt.Println(index, name)
		}()
	}
	//
	for index, name := range mySlice {
		go func(index int, name string) {
			fmt.Println(index, name)
		}(index, name)
	}
	time.Sleep(1 * time.Second)
	// 3 Hossein
	// 0 Ali
	// 1 Reza
	// 2 Amir
}
