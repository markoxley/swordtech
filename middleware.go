package swordtech

import (
	"net/http"
)

var middleware func(w http.ResponseWriter, r *http.Request) bool
var errorMiddleware func(w http.ResponseWriter, r *http.Request, err error) bool

// middlewareHandler is the middleware method for all requests.
// This function adds logging and testing for user signed in
func middlewareHandler(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w2 := NewResponse(w)
		r2 := NewRequest(r)

		// Defer logging so every request is logged
		defer func() {
			Log(&r2.Method, &r2.IPAddress, r2.RequestURI, w2.Status())
		}()

		path := r2.URL.Path

		// If the request if for a static file, serve the request
		if len(path) > 4 {
			if path[:5] == "/img/" || path[:5] == "/css/" || path[:5] == "/fnt/" || path[:4] == "/js/" {
				n.ServeHTTP(w2, r)
				return
			}
		}

		// Get the session for the user
		_, err := Session(r)

		// If the session could not be returned, or created, an error has occured, so return the 500 error
		if err != nil {
			doError := errorMiddleware == nil
			if !doError {
				doError = !errorMiddleware(w2, r, err)
			}
			if doError {
				http.Error(w2, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if middleware != nil && middleware(w2, r) {
			return
		}
		n.ServeHTTP(w2, r)
	})
}

// SetMiddleware sets the custom middleware
// Return true if the middleware handles the request
func SetMiddleware(f func(w http.ResponseWriter, r *http.Request) bool) {
	middleware = f
}

// SetErrorMiddleware sets the custom error handler
// Return true if the error middleware handles the request
func SetErrorMiddleware(f func(w http.ResponseWriter, r *http.Request, err error) bool) {
	errorMiddleware = f
}
