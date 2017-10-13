package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/patientplatypus/gorest/config"

	"github.com/gorilla/handlers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "patientplatypus"
	password = "Fvnjty0b"
	dbname   = "dungeon_world"
)

// var DB = config.Sql_connect()
var DB config.Sqlconfig

func main() {
	DB = config.Sql_config()
	router := NewRouter()
	os.Setenv("ORIGIN_ALLOWED", "*")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}
