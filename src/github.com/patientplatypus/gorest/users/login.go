package users

import (
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func (loginjson *User) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := config.KeyStore().Get(r, "cookie-name")
	log.Print("After assigning session in login")
	log.Print("loginjson: ", loginjson)
	log.Print("&loginjson.Username: ", &loginjson.Username)
	log.Print("loginjson.Username: ", loginjson.Username)
	// marshalledusername, error := json.Marshal(loginjson.Username)
	// if error != nil {
	// 	panic(error)
	// }
	var tempvar string

	config.DB.QueryRow("SELECT username FROM users WHERE username=$1;", loginjson.Username).Scan(tempvar)
	// log.Print("err: ", err)
	log.Print("got to here")
	var err string
	if err == "nil" {
		// 1 row
		log.Print("Found username")
		var passwordindatabase string
		config.DB.QueryRow("SELECT password FROM users WHERE username=$1;", &loginjson.Username).Scan(&passwordindatabase)
		if passwordindatabase == loginjson.Password {
			log.Print("username and password match!")
			session.Values["authenticated"] = true
			session.Values["username"] = &loginjson.Username
			config.KeyStore().Save(r, w, session)
		} else {
			log.Print("username and password don't match!")
			session.Values["authenticated"] = false
			session.Values["username"] = ""
			config.KeyStore().Save(r, w, session)
		}
	} else {
		//empty result or error
		log.Print("Username not found or there was an error: ", err)
	}
}
