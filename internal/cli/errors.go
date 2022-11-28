package cli

import "errors"

const (
	UserKind = iota
	SystemKind
)

var (
	ErrServiceAlreadyExists = NewErrorID("errServiceAlreadyExists", errors.New("path already exists"), UserKind)
	ErrServiceInternal      = NewErrorID("errServiceInternal", errors.New("internal error"), SystemKind)
)

type ErrorID struct {
	ID       string
	Err      error
	Kind     int
	MsgError error
}

func NewErrorID(id string, err error, kind int) *ErrorID {
	return &ErrorID{
		ID:   id,
		Err:  err,
		Kind: kind,
	}
}

func (e *ErrorID) Error() string {
	return e.Err.Error()
}

func (e *ErrorID) Unwrap() error {
	return e.Err
}

func (e *ErrorID) Wrap(err error) *ErrorID {
	e.MsgError = err
	return e
}
