package validation

import "net/http"

// Ruler is the interface for validation rules
type Ruler interface {
	validate(r *http.Request) bool
	Failed() bool
	ErrorMessage() string
	WebError() *webError
}
