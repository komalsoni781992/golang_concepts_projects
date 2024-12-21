package main

import (
	"fmt"
)

func main() {
	// if we don't provide any values to the variable, they get initialized with there default values
	// var block
	var (
		name  = "ajay"
		age   = 30
		marks = 80.5
	)
	var (
		a string
		b int
		c uint8
	)

	/*const (
		d string  // const values have to be initialized
	)*/
	const key = "some key"
	fmt.Println(name, age, marks, key)
	fmt.Println(a, b, c)

	// peek into it for design pattern
	//os.OpenFile()
	//os.O_RDONLY
	//time.Second
	//http.StatusNotFound

	//log.New(os.Stdout, log.LstdFlags)
	//l.Println("hello")
	//
	//math.MaxUint8
}
