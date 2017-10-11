package users

import (
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	store := config.KeyStore()
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
