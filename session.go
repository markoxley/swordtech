package swordtech

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// SessionName is the name of the session cookie
	sessionName string
	// Store is the main store for the sessions
	store *sessions.CookieStore
)

// CcnfigureSession configures the sessions
func configureSession(name string, key string) {
	sessionName = name
	store = sessions.NewCookieStore([]byte(name), []byte(key))
	store.Options.Secure = true
}

// Session returns the current session
func Session(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, sessionName)
}

// SessionName returns the name of the session
func SessionName() string {
	return sessionName
}

// GetSessionVar returns the value of the specified session variable
func GetSessionVar(r *http.Request, name string) interface{} {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return nil
	}

	return session.Values[name]
}

// SetSessionVar sets and saves the session variable
func SetSessionVar(r *http.Request, w http.ResponseWriter, name string, value interface{}) {
	if session, err := store.Get(r, sessionName); err == nil {
		session.Values[name] = value
		session.Save(r, w)
	}
}

// RemoveSessionVar removes the session variable
func RemoveSessionVar(r *http.Request, w http.ResponseWriter, name string) {
	if session, err := store.Get(r, sessionName); err == nil {
		delete(session.Values, name)
		session.Save(r, w)
	}
}

// RemoveAllSessionVars clears all the session variables
func RemoveAllSessionVars(r *http.Request, w http.ResponseWriter) {
	if session, err := store.Get(r, sessionName); err == nil {
		for k := range session.Values {
			delete(session.Values, k)
		}
		session.Save(r, w)
	}
}
