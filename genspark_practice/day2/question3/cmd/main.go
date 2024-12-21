package main

import (
	"fmt"
	"question3/auth"
)

// go build ./cmd  to create cmd.exe in module folder assumes it need to build whole project otherwise it can build single file also
// go run ./cmd/.
func main() {
	PrintSetup()
	auth.Authenticate("Komal")
	fmt.Print("end of main")
}
