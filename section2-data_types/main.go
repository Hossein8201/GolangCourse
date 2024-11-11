package main

import (
	"enum_const/enum_and_const"
	"enum_const/pointer"
	"enum_const/rune"
	"fmt"
	"math/bits"
	"unsafe"
)

type example struct {
	name string
	age  int
	flag bool
}

func main() {
	var num1 int
	var num2 int8
	var num3 int16
	var num4 int32
	var num5 int64
	var num6 = bits.UintSize
	var float1 float32
	var float2 float64
	//You can define variable the primitive types:
	var _ string
	var _ bool
	var _ int
	fmt.Printf("int has %d bytes\n", unsafe.Sizeof(num1))
	fmt.Printf("int8 has %d bytes\n", unsafe.Sizeof(num2))
	fmt.Printf("int16 has %d bytes\n", unsafe.Sizeof(num3))
	fmt.Printf("int32 has %d bytes\n", unsafe.Sizeof(num4))
	fmt.Printf("int64 has %d bytes\n", unsafe.Sizeof(num5))
	fmt.Println(num6)
	fmt.Printf("float32 has %d bytes\n", unsafe.Sizeof(float1))
	fmt.Printf("float64 has %d bytes\n", unsafe.Sizeof(float2))

	arr1 := [...]int{1, 2, 3, 4}
	arr2 := arr1
	arr3 := &arr1
	arr3[1] = 5
	fmt.Print("arr1 is : ")
	fmt.Println(arr1)
	fmt.Print("arr2 is : ")
	fmt.Println(arr2)
	fmt.Print("arr3 is : ")
	fmt.Println(arr3) // &[1 5 3 4]

	slice1 := arr2[1:3]
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Print("slice1 is : ")
	fmt.Println(slice1)
	fmt.Print("map1 is : ")
	fmt.Println(map1) // map[a:1 b:2 c:3]

	example1 := example{
		name: "test",
		age:  10,
		flag: true,
	}
	example2 := example1
	example2.flag = false
	example2.name = "answer"
	fmt.Print("example1 is : ")
	fmt.Println(example1)
	fmt.Print("example2 is : ")
	fmt.Println(example2)
	// Enum and Const variables
	var output *map[string]int
	output = enum_and_const.GiveInformation()
	println(*output)
	// Pointer
	pointer.ExamplePointer()
	// Rune types
	rune.ExampleRune()
}
