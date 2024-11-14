package __SwitchCase

import "fmt"

func SwitchCaseStatement() {
	var score float32
	fmt.Scanln(&score)
	//
	fmt.Println("the first type")
	switch score {
	case 20, 19, 18, 17:
		fmt.Println("Very good")
		fmt.Println("A")
	case 16, 15:
		fmt.Println("good")
		fmt.Println("B")
	case 14, 13, 12:
		fmt.Println("so so")
		fmt.Println("C")
	case 11, 10:
		fmt.Println("bad")
		fmt.Println("D")
	default:
		fmt.Println("fall")
		fmt.Println("F")
	}
	//
	fmt.Println("the second type")
	switch {
	case score >= 17:
		fmt.Println("Very good")
		fmt.Println("A")
	case score >= 15 && score < 17:
		fmt.Println("good")
		fmt.Println("B")
	case score >= 12 && score < 15:
		fmt.Println("so so")
		fmt.Println("C")
	case score >= 10 && score < 12:
		fmt.Println("bad")
		fmt.Println("D")
	default:
		fmt.Println("fall")
		fmt.Println("F")
	}
}
