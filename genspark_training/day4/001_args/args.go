package main

import (
	"fmt"
	"os"
)

// First arg is location of temorary .exe file - C:\Users\zubin\AppData\Local\Temp\go-build3384109166\b001\exe\args.exe
func main() {
	fmt.Println(os.Args)
}
