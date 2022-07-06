package validation

import (
	"fmt"
	"net/http"
)

// RequiredRule tests that the value of the field is not empty
type RequiredRule struct {
	baseRule
}

func (rq *RequiredRule) validate(r *http.Request) bool {
	tgt := r.PostFormValue(rq.fieldName)
	if len(tgt) == 0 {
		rq.fail = true
		*rq.error = fmt.Sprintf("%s is a required field", rq.fieldName)
	}
	return !rq.fail
}

// Required creates a new integer rule. This rule only requires the field name to be passed
func Required(fieldName string) *RequiredRule {
	return &RequiredRule{
		baseRule: baseRule{
			fieldName: fieldName,
		},
	}
}
