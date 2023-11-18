package errors

type TemplateParsingError struct {
	Template string
}

func NewTemplateParsingError(template string) *TemplateParsingError {
	return &TemplateParsingError{Template: template}
}

func (e *TemplateParsingError) Error() string {
	return "Error parsing template: " + e.Template
}
