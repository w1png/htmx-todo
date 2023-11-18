package errors

type EnvironmentVariableNotFoundError struct {
	VariableName string
}

func NewEnvironmentVariableNotFoundError(variableName string) *EnvironmentVariableNotFoundError {
	return &EnvironmentVariableNotFoundError{VariableName: variableName}
}

func (e *EnvironmentVariableNotFoundError) Error() string {
	return "Environment variable not found: " + e.VariableName
}
