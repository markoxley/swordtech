package validation

import (
	"fmt"
	"net/http"
)

// MatchRule is the rule used for testing two fields match
type MatchRule struct {
	baseRule
	field string
}

func (m *MatchRule) validate(r *http.Request) bool {
	tgt := r.PostFormValue(m.fieldName)
	oth := r.PostFormValue(m.field)
	if tgt != oth {
		m.fail = true
		*m.error = fmt.Sprintf("%s and %s must match", m.fieldName, m.field)
	}
	return !m.fail
}

// Match creates a new match rule. This rule requires the name of the other field to be passed
func Match(fieldName string, field string) *MatchRule {
	return &MatchRule{
		baseRule: baseRule{
			fieldName: fieldName,
		},
		field: field,
	}
}
