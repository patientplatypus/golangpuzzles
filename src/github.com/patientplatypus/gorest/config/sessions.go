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

func KeyStore() (store Store) {
	log.Print("inside KeyStore")
	key := []byte("H(@>gCYB}xU@f{Q?yV{9F2PM~O1vy=")
	store = sessions.NewCookieStore(key)
	return store
}
