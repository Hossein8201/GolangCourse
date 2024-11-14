package __If_Else

import "fmt"

func IfElseStatement() {
	var salary int
	fmt.Scanln(&salary)
	if salary < 1000 {
		fmt.Println("No")
	} else if salary == 1000 {
		fmt.Println("Unknown")
	} else {
		fmt.Println("Yes")
	}
}
