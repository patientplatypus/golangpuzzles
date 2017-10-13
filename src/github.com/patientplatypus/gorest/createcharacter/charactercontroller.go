package createcharacter

import (
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

var Username string
var Checkok bool

func SessionsCheck(w http.ResponseWriter, r *http.Request) (username string, checkok bool) {
	store := config.KeyStore()
	session, _ := store.Get(r, "cookie-name")
	log.Print("username: ", session.Values["username"])
	log.Print("authenticated: ", session.Values["authenticated"])
	if session.Values["username"] == nil {
		if session.Values["authenticated"] == false {
			log.Print("Verboten!")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return "nil", false
		}
	}
	return session.Values["username"].(string), true
}

func NewCharacter(w http.ResponseWriter, r *http.Request) {
	Username, Checkok = SessionsCheck(w, r)
	charactername := r.FormValue("charactername")
	log.Print("username: ", Username)
	if Checkok == false {
		return
	}
	RegisterCharacter(charactername)

}
