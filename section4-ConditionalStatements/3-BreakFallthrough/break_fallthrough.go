package __BreakFallthrough

import "fmt"

func BreakFallthroughStatement() {
	var num int
	fmt.Scanf("%d", &num)
	//
	fmt.Println("the break statement")
	switch num {
	case 1:
		fmt.Print("1")
		break
		fmt.Print(" the number")
	case 2:
		fmt.Print("2")
		fmt.Print(" the number")
	}
	//
	fmt.Println("\nthe fallthrough statement")
	switch num {
	case 1:
		fmt.Print("1")
		fmt.Print(" the number")
		fallthrough
	case 2:
		fmt.Print(" 2")
		fmt.Print(" the number")
	}
}
