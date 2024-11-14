package __Consts

import "fmt"

func Constants() {
	const pi float64 = 3.1415926
	const name string = "Hossein"
	fmt.Printf("%s is constant value: %f\n", name, pi)
	const (
		num1   = 12.55
		num2   = 13.1415926
		letter = 'd'
	)
	fmt.Println(num1, num2, letter)
	const urlGoogle = "https://google"
	const urlMap = ".com"
	println(urlGoogle + urlMap)
}
