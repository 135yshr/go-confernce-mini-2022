package usecase

import (
	"fmt"
)

type UseCaseError struct {
	code    int
	message string
}

func (e *UseCaseError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.code, e.message)
}

func NewError(code int, message string) *UseCaseError {
	return nil
}
