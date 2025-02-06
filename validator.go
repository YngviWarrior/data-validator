package validator

type validation struct {
	Error []string
}

type ValidationInterface interface {
	Validation(o any) []string
}

func NewValidator() ValidationInterface {
	return &validation{}
}
