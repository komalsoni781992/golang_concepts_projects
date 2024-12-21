package main

import (
	"fmt"
)

func main() {
	// make preallocates a backing array
	// very helpful if we have idea about the number of request ,
	//we can create size of array according to that
	// in this case append doesn't have to allocate the memory each time
	i := make([]int, 0, 50) // type, len, cap
	Inspect("i", i)
	i = append(i, 100)
	fmt.Println(i)
}

func Inspect(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("memory address %p\n", slice)
	fmt.Println(slice)
	fmt.Println()
}
