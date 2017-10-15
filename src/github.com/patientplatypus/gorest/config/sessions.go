package config

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type Options struct {
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HttpOnly bool
}

type Session struct {
	ID      string
	Values  map[interface{}]interface{}
	Options *Options
	IsNew   bool
}

type Store interface {
	Get(r *http.Request, name string) (*sessions.Session, error)
	New(r *http.Request, name string) (*sessions.Session, error)
	Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error
}

var SessionsStore = sessions.NewCookieStore([]byte("H(@>gCYB}xU@f{Q?yV{9F2PM~O1vy="))

func init() {
	SessionsStore.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}
}

func KeyStore() (store Store) {

	log.Print("inside KeyStore")
	// log.Print("session.Value['username'] before key,store: ", session.Value["username"])
	// key := []byte("H(@>gCYB}xU@f{Q?yV{9F2PM~O1vy=")
	// store = sessions.NewCookieStore(key)
	// log.Print("session.Value['username'] after key, store: ", session.Value["username"])
	store = SessionsStore
	log.Print("Value of store is : ", store)
	// log.Print("Value of store.Values['username']: ", store.Get(r, "cookie-name"))
	return store
}
