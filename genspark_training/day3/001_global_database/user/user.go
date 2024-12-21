package user

import (
	"fmt"
	"globaldatabase/database"
)

func AddUser(db database.SQLDb) {
	fmt.Println("adding to db ", db)
}
