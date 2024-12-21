package main

import (
	"globaldatabase/database"
	"globaldatabase/user"
)

func main() {
	db := database.InitDB("localhost:3306:postgres")
	//database.DB = "random" // anyone can change the value of the DB here
	//var i string = "SQL" - cannot be passed in AddUser
	user.AddUser(db)
}
