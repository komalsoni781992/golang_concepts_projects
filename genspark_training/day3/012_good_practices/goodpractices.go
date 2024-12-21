package main

import (
	"fmt"
)

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3)
	x = update(x)

	fmt.Println(x)
}

func update(s []int) []int {
	// if you want to update the value , then passing the reference of existing backing array is fine
	s[0] = 100

	// if you want to append
	// if you are appending the values to the slice, always return the slice
	// returning the slice would ensure that caller has updated reference
	s = append(s, 111, 222, 333)

	Inspect("after s", s)

	//and slice is not returned back, then code review should stop
	return s

}

func Inspect(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("memory address %p\n", slice)
	fmt.Println(slice)
	fmt.Println()
}

// slices and arrays are passed as value
