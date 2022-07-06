package validation

import (
	"fmt"
	"net/http"
	"regexp"
)

// PasswordRule is the rule used for testing the complexity of passwords
type PasswordRule struct {
	baseRule
}

func (p *PasswordRule) validate(r *http.Request) bool {
	var re *regexp.Regexp
	tgt := r.PostFormValue(p.fieldName)
	ptns := []string{
		`(?m)([!@£#$%^&*\(\)\-=_+\[\]{};:,.<>])`,
		`([a-z])`,
		`([A-Z])`,
		`([0-9])`,
	}
	found := 0
	for _, ptn := range ptns {
		re = regexp.MustCompile(ptn)
		if len(re.FindStringIndex(tgt)) > 0 {
			found++
		}
	}
	if found < 3 || len(tgt) < 8 {
		p.fail = true
		*p.error = fmt.Sprintf("%s is not secure enough", p.fieldName)

	}
	return !p.fail
}

// Password creates a new password rule. This rule only requires the field name to be passed
func Password(fieldName string) *PasswordRule {
	return &PasswordRule{
		baseRule: baseRule{
			fieldName: fieldName,
		},
	}
}
