package error_handling

import "fmt"

type ErrNotFound struct {
	err error
}

func NewErrNotFound(err error) *ErrNotFound {
	return &ErrNotFound{err: err}
}

func (e *ErrNotFound) Error() string {
	return fmt.Errorf("not found: %w", e.err).Error()
}

func (e *ErrNotFound) Unwrap() error {
	return e.err
}
