package users

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/patientplatypus/gorest/config"
)

func (user *User) Create() {
	if err := config.DB.QueryRow("SELECT username FROM users WHERE username=$1;", &user.Username).Scan(&user.Username); err == nil {
		// 1 row
		fmt.Println("there was 1 or more rows returned and err is: ", err)
		fmt.Println("You cannot add a user with that name as a user with that name already exists! oh no!")
		QueryAll()
	} else if err == sql.ErrNoRows {
		// empty result
		fmt.Println("no rows from sql and err is: ", err)
		config.DB.QueryRow("insert into users (username, password, id) values ($1, $2, $3);", user.Username, user.Password, rand.Intn(999999))
		fmt.Println("User inserted and rows is equal to: ", err)
	} else {
		// error
		fmt.Println("Something has gone wrong. err is: ", err)
	}
}
