package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)

	}

	for i, j := 0, 10; i <= 10; i++ {
		fmt.Println(j)
	}

	i := 0
	for ; i <= 10; i++ {

	}

	i = 0
	for i <= 10 {
		//do work here

		i++

	}

	// ranges will be covered later

}
