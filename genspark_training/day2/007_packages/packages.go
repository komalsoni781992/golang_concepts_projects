package main

// If I rename this package name to calc then it will be different from other calc due to different directory structure

import (
	c "crypto/rand"
	"fmt"
	"math/rand"     // If 2 packages have same name collision has to be handled by package renaming of at least 1 package
	"packages/calc" // moduleName/packageName
	"packages/some_pack"
)

// package main is also known as a binary package,
// this is the only package that can run

// func main is required to run a binary package
// https://google.github.io/styleguide/go/best-practices#util-packages
func main() {
	fmt.Println("running the app")
	//x:= fmt.Sprintln("Hello") - returns the string that has been printed. It uses several unexported functions
	calc.Add(1, 2)
	// calc.subtract(2, 44) - cannot be called as it has not been exported

	some_pack.NotRecommended() // don't use _ in package names
	rand.Int()
	c.Int(nil, nil)

	// don't use util packages in naming
	//ioutil.ReadFile()
	//ioutil.ReadAll()
	//os.ReadFile()
	//io.ReadAll()

	// Avoid util packages
}
