package werror

import (
	"errors"
)

// ErrorTemplate is a template error message, used to create comparible Errors
type ErrorTemplate interface {
	Here(msgs ...string) *Error
	Is(err error) bool
}

type errorTemplate struct {
	err error
	// defaultCode int // TODO
}

// Template creates a new ErrorTemplate, use Here() to get an error
func Template(msg string) ErrorTemplate {
	return &errorTemplate{
		err: errors.New(msg),
	}
}

// Here creates the error and records a stacktrace with optional tags
func (t errorTemplate) Here(msgs ...string) *Error {

	werr := &Error{
		err: t.err,
	}
	werr.addStack(1)

	werr = Wrap(werr, msgs...)

	return werr
}

// Is returns true if the error matches this Template
func (t *errorTemplate) Is(err error) bool {
	return Is(err, t.err)
}
