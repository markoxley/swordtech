package validation

import (
	"fmt"
	"net/http"
	"strconv"
)

// IntegerRule is the rule used for testing for integer numbers
type IntegerRule struct {
	baseRule
}

func (i *IntegerRule) validate(r *http.Request) bool {
	tgt := r.PostFormValue(i.fieldName)

	if _, err := strconv.ParseInt(tgt, 10, 64); err != nil {
		i.fail = true
		*i.error = fmt.Sprintf("%s must be whole point number", i.fieldName)
	}
	return !i.fail
}

// Integer creates a new integer rule. This rule only requires the field name to be passed
func Integer(fieldName string) *IntegerRule {
	return &IntegerRule{
		baseRule: baseRule{
			fieldName: fieldName,
		},
	}
}
