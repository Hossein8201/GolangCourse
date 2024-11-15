package main

import (
	arrays "CompositeDataTypes/1-Arrays"
	slices "CompositeDataTypes/2-Slices"
	moreOfSlices "CompositeDataTypes/3-MoreOfSlices"
	maps "CompositeDataTypes/4-Maps"
	"fmt"
)

func main() {
	arrays.ArraysExample()
	fmt.Println()
	slices.SliceExample()
	fmt.Println()
	moreOfSlices.MoreOfSlices()
	fmt.Println()
	maps.MapExample()
}
