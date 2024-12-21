package user

import (
	"fmt"
	//"question2/auth" - creates import cycle, it is not allowed in Go
	/***How to solve the import cycle**
	  Extract the common piece of functionality in a separate package
	  Import the new package where functionality is needed*/)

var userName string

// AddToDb - authenticates name and add to DB
func AddToDb(name string) string {
	userName = name
	fmt.Println(userName)
	//auth.Authenticate(name)
	return name
}
