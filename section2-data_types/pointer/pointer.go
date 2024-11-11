package pointer

import (
	"fmt"
)

func ExamplePointer() {
	i, j := 1, 2
	var pi, pj *int = &i, &j
	fmt.Printf("the address of i is %d == %d\n", pi, &i)
	fmt.Printf("the address of j is %d == %d\n", pj, &j)
}
