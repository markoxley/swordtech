package validation

import "net/http"

// Validate tests the passed rules against the http.Request
func Validate(r *http.Request, rules ...Ruler) bool {
	ok := true
	for _, rule := range rules {
		if !rule.validate(r) {
			ok = false
		}
	}

	return ok
}
