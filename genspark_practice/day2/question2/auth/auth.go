package auth

import (
	"fmt"
	"question2/user"
)

// Authenticate authenticates user
func Authenticate(user string) {
	fmt.Println("Authenticating user: ", user)
}

// Name adds user to db
func Name() {
	fmt.Println(user.AddToDb("Komal"))
}
