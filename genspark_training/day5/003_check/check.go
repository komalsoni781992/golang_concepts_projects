package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(10, 20)) // callign named function

	func(x, y int) int {
		return x + y
	}(10, 20) // () this is to invoke the anonymous function

	//fmt.Println(greet)
	fmt.Println(call_greet(greet_two, "Komal"))
	fmt.Println(call_greet(greet_one(), "Komal"))
	fmt.Println(call_greet(greet_three, "Komal"))

}

func call_greet(fn func() string, name string) string {
	x := fn() + name
	return x
}

// greet can return itself leading to stack overflow
func greet_one() func() string {
	return func() string {
		return "hello"
	}
}

func greet_two() string {
	return "hello"
}

func greet_three() string {
	return func() string {
		return "hello"
	}()
}
