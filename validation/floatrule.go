package validation

import (
	"fmt"
	"net/http"
	"strconv"
)

// FloatRule is the rule used for testing for floating point numbers
type FloatRule struct {
	baseRule
}

func (f *FloatRule) validate(r *http.Request) bool {
	tgt := r.PostFormValue(f.fieldName)

	if _, err := strconv.ParseFloat(tgt, 64); err != nil {
		f.fail = true
		*f.error = fmt.Sprintf("%s must be floating point number", f.fieldName)
	}
	return !f.fail
}

// Float creates a new float rule. This rule only requires the field name to be passed
func Float(fieldName string) *FloatRule {
	return &FloatRule{
		baseRule: baseRule{
			fieldName: fieldName,
		},
	}
}
