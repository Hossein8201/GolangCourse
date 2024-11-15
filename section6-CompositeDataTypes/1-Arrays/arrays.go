package __Arrays

import "fmt"

func ArraysExample() {
	var myArray1 [2]int // initialized to zero
	var myArray2 [5]int = [5]int{1, 2, 3}
	var myArray3 = [6]int{1, 2, 3}
	myArray4 := [7]int{1, 2, 3}
	myArray5 := [...]int{1, 2, 3}
	//
	fmt.Println("the members :")
	fmt.Println(myArray1)
	fmt.Println(myArray2)
	fmt.Println(myArray3)
	fmt.Println(myArray4)
	fmt.Println(myArray5)
	//
	fmt.Println("the size of that")
	fmt.Println(len(myArray1))
	fmt.Println(len(myArray2))
	fmt.Println(len(myArray3))
	fmt.Println(len(myArray4))
	fmt.Println(len(myArray5))
	//
	myArray1[1] = 4
	fmt.Println("the new array", myArray1)
	//
	intArray1 := [...]int{1, 2, 3}
	intArray2 := intArray1
	intArray3 := &intArray1 // the type of this is a Pointer
	intArray1[0] = -1
	fmt.Println("the address of them is : ")
	fmt.Printf("%p\n", &intArray1)
	println(&intArray1)
	fmt.Printf("%p\n", &intArray2)
	println(&intArray2)
	fmt.Println(&intArray3)
	fmt.Println("the new array", intArray1)
	fmt.Println("the copy array", intArray2)
	fmt.Println("the copy with '&' array", intArray3)
}
