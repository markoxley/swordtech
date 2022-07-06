package validation

import (
	"fmt"
	"net/http"
)

// MinRule is the rule used for testing tthe field has the minimum number of characters
type MinRule struct {
	baseRule
	value int
}

func (m *MinRule) validate(r *http.Request) bool {
	tgt := r.PostFormValue(m.fieldName)
	if len(tgt) < m.value {
		m.fail = true
		*m.error = fmt.Sprintf("%s must be at least %d characters long", m.fieldName, m.value)
	}
	return !m.fail
}

// Min creates a new min rule. This rule requires the minimum number of characters
func Min(fieldName string, length int) *MinRule {
	return &MinRule{
		baseRule: baseRule{
			fieldName: fieldName,
		},
		value: length,
	}
}
