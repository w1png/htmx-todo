package errors

type InvalidStorageTypeError struct {
	StorageType string
}

func (e *InvalidStorageTypeError) Error() string {
	return "Invalid storage type: " + e.StorageType
}

func NewInvalidStorageTypeError(storageType string) *InvalidStorageTypeError {
	return &InvalidStorageTypeError{StorageType: storageType}
}

type NotFoundError struct {
	Object string
}

func (e *NotFoundError) Error() string {
	return "Not found: " + e.Object
}

func NewNotFoundError(object string) *NotFoundError {
	return &NotFoundError{Object: object}
}
