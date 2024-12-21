package main

import (
	"fmt"
	"strconv"
)

func main() {
	if v, err := ParseStringToFloat64("123.55"); err == nil {
		fmt.Println(v)
	} else {
		panic("Invalid input")
	}
}

// ParseStringToFloat64 - parese string to float64
func ParseStringToFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return f, nil
	}
	return 0, err
}
