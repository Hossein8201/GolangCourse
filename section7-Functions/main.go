package main

import (
	inputsoutputs "Functions/1-InputsAndOutputs"
	anonymousfunctions "Functions/2-AnonymousFunctions"
	closures "Functions/3-Closures"
	__Defer "Functions/4-Defer"
	"fmt"
)

func main() {
	fmt.Print("Print1 : ")
	inputsoutputs.Print1()
	//
	fmt.Println("Print2 :", inputsoutputs.Print2())
	//
	str3 := "Print3 : hello world"
	fmt.Println(inputsoutputs.Print3(str3))
	//
	str4 := "Print4 : hello world"
	number4 := 4
	float4, str4 := inputsoutputs.Print4(number4, str4)
	fmt.Println(float4, str4)
	//
	str5 := "Print5 : hello world"
	number5 := 5
	float5, str5 := inputsoutputs.Print5(number5, 2, str5)
	fmt.Println(float5, str5)
	//
	slice6 := []int{1, 2, 3, 4, 5}
	fmt.Println(inputsoutputs.Print6(slice6...)) // Unpack slice
	fmt.Println(inputsoutputs.Print6(6, 7, 8, 9, 10))
	//
	inputsoutputs.Print7(1, "you", false, "are", 3, true, 12.666, "Hossein")
	// Anonymous function
	fmt.Println()
	anonymousfunctions.Example()
	// Closures function
	fmt.Println()
	closures.Example()
	//
	fmt.Println()
	__Defer.Defer()
	destination := "D:\\__desktop folders\\Golang Course\\section7-Functions\\4-Defer\\destination.txt"
	start := "D:\\__desktop folders\\Golang Course\\section7-Functions\\4-Defer\\start.txt"
	fmt.Println(__Defer.CopyFile(destination, start))
	//
}
