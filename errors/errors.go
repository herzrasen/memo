package errors

type InvalidKeyError struct {}

func (e *InvalidKeyError) Error() string {
	return "invalid key provided"
}

func NewInvalidKeyError() *InvalidKeyError {
	return &InvalidKeyError{}
}

