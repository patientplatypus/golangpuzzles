package config

import "fmt"

type Sqlconfig struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

// func (post *Post) Create() (err error) {
// 	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.Id)
// 	return
// }

func Sql_config() (sqlconfig *Sqlconfig) {
	// panic("it works!")
	fmt.Println("Inside Sql_config")
	returnsql := Sqlconfig{
		host:     "localhost",
		port:     5432,
		user:     "patientplatypus",
		password: "Fvnjty0b",
		dbname:   "dungeon_world"}

	fmt.Println("Value of returnsql: ", returnsql)

	return &returnsql
}
