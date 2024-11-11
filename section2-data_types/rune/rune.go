package rune

import (
	"fmt"
	"unsafe"
)

func ExampleRune() {
	// common character or "ASCII"
	str1 := "hello world *"
	runes := []rune(str1)
	fmt.Println(runes)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%d: %v\n", i, runes[i])
	}
	// special character
	str2 := "سلام دنیا"
	fmt.Println(unsafe.Sizeof(str2)) // 16
	fmt.Println(len(str2))           // 17
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%d: %c\n", i, str2[i])
	}
	runes2 := []rune(str2)
	fmt.Println(runes2)      // [1587 1604 1575 1605 32 1583 1606 1740 1575]
	fmt.Println(len(runes2)) // 9
	for i := 0; i < len(runes2); i++ {
		fmt.Printf("%d: %c\n", i, runes2[i])
	}
}
