package errors

import (
	"github.com/pkg/errors"
)

// Application errors
const (
	NoType       = ErrorType(iota)
	ErrInvalid   = ErrorType(iota)
	ErrForbidden = ErrorType(iota)
	ErrNotFound  = ErrorType(iota)
	ErrInternal  = ErrorType(iota)
)

type ErrorType uint

type appError struct {
	errType       ErrorType
	originalError error
}

func (err appError) Error() string {
	return err.originalError.Error()
}

func (t ErrorType) New(msg string) error {
	return appError{errType: t, originalError: errors.New(msg)}
}

func (t ErrorType) Wrap(err error, msg string) error {
	return t.Wrapf(err, msg)
}

func (t ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, args...)

	return appError{errType: t, originalError: newErr}
}

func (t ErrorType) Error() error {
	return appError{errType: t}
}

func GetType(err error) ErrorType {
	if customErr, ok := err.(appError); ok {
		return customErr.errType
	}

	return NoType
}
