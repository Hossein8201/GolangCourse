package __defineVariables

import "fmt"

func DefineVariables() {
	//	1 - var varName type = value
	var number1 int = 1
	var name1 string = "hossein tatar"
	fmt.Println(number1, name1)
	// 2 - var varName type
	var number2 int
	var name2 string
	fmt.Println(number2, name2)
	// 3 - var varName = value
	var number3 = 12
	var number4 = 23.55
	var name3 = "Ali"
	fmt.Println(number3, number4, name3)
	// 4 - varName := value
	number5 := 55
	name4 := "Emad"
	fmt.Println(number5, name4)
	// grouping
	var num1, fNum, country = 55, 22.55, "Iran"
	num2, fNum2, country2 := 55, 22.55, "Iran"
	fmt.Println(num1, fNum, country)
	fmt.Println(num2, fNum2, country2)
	var (
		num3     = 33
		fNum3    = 44.44
		country3 = "America"
	)
	fmt.Println(num3, fNum3, country3)
}
