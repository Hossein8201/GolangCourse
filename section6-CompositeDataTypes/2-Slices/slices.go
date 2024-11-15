package __Slices

import "fmt"

func SliceExample() {
	myArr := [3]int{1, 2, 3}
	myArr2 := [...]int{1, 2, 3}
	mySlice := []int{1, 2, 3}
	fmt.Println(myArr, myArr2, mySlice)
	// use make :
	mySlice2 := make([]int, 10) // create a slice with len and cap 10
	fmt.Println("the len of mySlice2 is ", len(mySlice2), "and the cap of mySlice2 is ", cap(mySlice2))
	mySlice3 := make([]int, 10, 20) // create a slice with len 10 and cap 20
	fmt.Println("the len of mySlice3 is ", len(mySlice3), "and the cap of mySlice3 is ", cap(mySlice3))
	// create from Array
	mySlice4 := myArr[:]
	mySlice5 := myArr[1:]
	mySlice6 := myArr[:2]
	mySlice7 := myArr[1:2]
	fmt.Println(mySlice4, mySlice5, mySlice6, mySlice7)
	// changing
	myArr3 := [...]int{10, 20, 30, 40, 50}
	mySlice8 := myArr3[:]
	myArr3[0] = 100
	fmt.Println(myArr3, mySlice8) // [100 20 30 40 50] [100 20 30 40 50]
	// Appending
	mySlice8 = append(mySlice8, 60)
	mySlice8 = append(mySlice8, 70, 80, 90)
	myArr3[4] = 500
	fmt.Println(myArr3, mySlice8) // [100 20 30 40 500] [100 20 30 40 50 60 70 80 90]
}
