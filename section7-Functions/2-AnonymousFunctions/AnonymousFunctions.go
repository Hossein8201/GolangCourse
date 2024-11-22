package __AnonymousFunctions

import (
	"fmt"
	"sort"
)

func Example() {
	//
	func() {
		fmt.Println("func1")
	}()
	//
	myFunc := func() {
		fmt.Println("func2")
	}
	myFunc()
	myFunc()
	//
	fmt.Println(func(numbers ...int) (sum int) {
		for _, number := range numbers {
			sum += number
		}
		return
	}(1, 2, 3))
	//
	sumFunc := func(numbers ...int) (sum int) {
		for _, number := range numbers {
			sum += number
		}
		return
	}
	fmt.Println(sumFunc(1, 2, 3, 4, 5))
	fmt.Println(sumFunc(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	//
	mySlice := []int{4, 5, 2, 3, 7, 6, 8, 9, 10, 1}
	fmt.Println(mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i] < mySlice[j]
	})
	fmt.Println(mySlice)
	//
	mySlice = []int{4, 5, 2, 3, 7, 6, 8, 9, 10, 1}
	fmt.Println(mySlice)
	mySort := func(i, j int) bool {
		return mySlice[i] < mySlice[j]
	}
	sort.Slice(mySlice, mySort)
	fmt.Println(mySlice)
}
