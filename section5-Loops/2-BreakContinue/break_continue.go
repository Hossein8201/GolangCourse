package __BreakContinue

import (
	"fmt"
	"math/rand"
)

func BreakContinue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	//
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	//
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Int())
	//
	listName := []string{"ali", "reza", "hossein", "amir", "armin"}
	for _, name := range listName {
		fmt.Println(name)
	}
}
