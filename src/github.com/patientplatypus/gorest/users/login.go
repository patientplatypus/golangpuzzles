package users

import (
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func (user *User) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := config.KeyStore().Get(r, "cookie-name")
	if err := config.DB.QueryRow("SELECT username FROM users WHERE username=$1;", &user.Username).Scan(&user.Username); err == nil {
		// 1 row
		log.Print("Found username")
		var passwordindatabase string
		config.DB.QueryRow("SELECT password FROM users WHERE username=$1;", &user.Username).Scan(&passwordindatabase)
		if passwordindatabase == user.Password {
			log.Print("username and password match!")
			session.Values["authenticated"] = true
			session.Values["username"] = &user.Username
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
