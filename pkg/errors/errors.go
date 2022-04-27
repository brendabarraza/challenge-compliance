package errors

import (
	"errors"
)

var (
	// ErrEmptyResult ...
	ErrEmptyResult = errors.New("empty result")
	// ErrEmptyParameter ...
	ErrEmptyParameter = errors.New("empty parameter")
	// ErrNotDefined ...
	ErrNotDefined = errors.New("error not defined")
	// ErrBadConfiguration ...
	ErrBadConfiguration = errors.New("bad configuration")
	// ErrDriveClient...
	ErrDriveClient = errors.New("unable to retrieve Drive client")
	// ErrUnableFiles...
	ErrUnableFiles = errors.New("unable to retrieve files")
)
