package __MoreOfSlices

import "fmt"

func MoreOfSlices() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ChangeValueItem(mySlice) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println("after changing with items :", mySlice)
	ChangeValueIndex(mySlice) // [10 20 30 40 50 60 70 80 90 100]
	fmt.Println("after changing with index :", mySlice)
	AppendValueCall(mySlice) // [10 20 30 40 50 60 70 80 90 100]
	fmt.Println("after appending with call :", mySlice)
	AppendValueReference(&mySlice) // [10 20 30 40 50 60 70 80 90 100 110]
	fmt.Println("after appending with Reference :", mySlice)
	//
	var mySlice0 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	count := copy(mySlice0, mySlice[:11])
	fmt.Println("after copying :", mySlice0)                // [10 20 30 40 50 60 70 80 90 100]
	fmt.Println("the number of items that copied :", count) // 10
	//
	mySlice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	mySlice2 = append(mySlice2, mySlice...) // merge
	fmt.Println("after merging mySlice2 with mySlice :", mySlice2)
	//
	mySlice2 = mySlice2[:0]
	fmt.Println(mySlice2)
	fmt.Println("the len :", len(mySlice2))
	fmt.Println("the cap :", cap(mySlice2))
	mySlice = nil
	fmt.Println(mySlice)
	fmt.Println("the len :", len(mySlice))
	fmt.Println("the cap :", cap(mySlice))
}

func ChangeValueItem(mySlice []int) {
	for _, value := range mySlice {
		value = value * 10
	}
}

func ChangeValueIndex(mySlice []int) {
	for index, value := range mySlice {
		mySlice[index] = value * 10
	}
}

func AppendValueCall(mySlice []int) {
	mySlice = append(mySlice, 110)
}

func AppendValueReference(mySlice *[]int) {
	*mySlice = append(*mySlice, 110)
}
