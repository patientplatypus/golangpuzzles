package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/patientplatypus/gorest/config"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "patientplatypus"
	password = "Fvnjty0b"
	dbname   = "dungeon_world"
)

func main() {
	router := NewRouter()
	config.DB = config.Sql_connect()
	log.Fatal(http.ListenAndServe(":8080", router))
}
