package validation

type baseRule struct {
	fieldName string
	fail      bool
	error     *string
}

func (b *baseRule) Failed() bool {
	return b.fail
}

func (b *baseRule) ErrorMessage() string {
	if b.fail {
		if b.error != nil {
			return *b.error
		}
		return "Error"
	}
	return ""
}

func (b *baseRule) WebError() *webError {
	if b.Failed() {
		return &webError{
			FieldName: b.fieldName,
			Message:   b.ErrorMessage(),
		}
	}
	return nil
}
