package __InputsAndOutputs

import (
	"fmt"
)

func Print1() {
	fmt.Println("Print1 : hello world")
}

func Print2() string {
	return "Print2 : hello world"
}

func Print3(example string) string {
	return example
}

func Print4(number int, example string) (float64, string) {
	var float4 float64
	float4 = float64(number) * 1
	return float4, example
}

func Print5(number1, number2 int, example string) (float5 float64, str5 string) {
	float5 = float64(number1)*0.6 + float64(number2)
	str5 = example
	return
}

func Print6(slice6 ...int) (sum, mul int) {
	mul = 1
	for _, number := range slice6 {
		sum += number
		mul *= number
	}
	return
}

func Print7(slice7 ...interface{}) {
	for index, value := range slice7 {
		fmt.Printf("%d, %v\n", index, value)
	}
}
