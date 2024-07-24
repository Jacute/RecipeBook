package database

import "fmt"

type AlreadyExistsError struct {
	Message string
}

// Error возвращает строку с описанием ошибки.
func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("AlreadyExistsError: %s", e.Message)
}
