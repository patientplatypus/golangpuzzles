package users

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patientplatypus/gorest/config"
)

//read/write sql/db

// func (post *Post) Create() (err error) {
// 	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.Id)
// 	return
// }
//
// func init() {
// 	var err error
// 	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
// 	if err != nil {
// 		panic(err)
// 	}
// }

//hook to routes

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	class := r.FormValue("class")
	fmt.Fprintln(w, "Your Class: ", class)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sqlConfig := config.Sql_config()
	config.Sql_config()
	fmt.Fprintln(w, "Login: ", vars["username"], " Password: ", vars["password"], " Sql_config: ", sqlConfig)
}
