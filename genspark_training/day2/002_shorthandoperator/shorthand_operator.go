package main

import "fmt"

func main() {
	a, greet := 1, "hello" // we are creating two variables

	// var greet int = 10 - This won't work as greet is being redeclared

	// greet = 20 - This won't work due to type mismatch

	a, ok := 10, true // if a variable already exists, it would update the existing one,
	// and create other variables which are not already present

	// a := 20 // this would not work // we need at least one new variable on the left side

	fmt.Println(a, ok, greet)
}
