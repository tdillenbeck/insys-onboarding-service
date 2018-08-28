package werror

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// This is a magic error type, which records a stacktrace and optional Tags
type Error struct {
	err       error
	messages  []string
	clientMsg string
	stack     []uintptr
	tags      []Tags

	code Code

	// Reported records whether the message has been reported to Sentry.
	Reported bool
}

// Message returns the raw error message from the wrapped error, no tags
func (e *Error) Message() string {
	if e.clientMsg != "" {
		return e.clientMsg
	}

	return e.err.Error()
}

// ExtraMessages returns the extra messages added to the wrapped error
func (e *Error) ExtraMessages() []string {
	return e.messages
}

type Stringer interface {
	String() string
}

// Error implements error.Error(), returns the message followed by the tags, like:
// "error foo=hi err=true", values are formatted like %v.
func (e *Error) Error() string {
	if e == nil {
		return "*werr.Error is nil!"
	}
	parts := []string{}
	for _, m := range e.messages {
		parts = append(parts, m)
	}
	parts = append(parts, e.err.Error())
	parts = []string{strings.Join(parts, " | ")}

	for _, ts := range e.tags {
		for k, v := range ts {
			if sv, ok := v.(Stringer); ok {
				v = sv.String()
			}

			parts = append(parts, fmt.Sprintf(" %s=[%#v]", k, v))
		}
	}
	return strings.Join(parts, "")
}

// AddTags appends tags to the list of tags.
func (e *Error) AddTags(tags ...Tags) *Error {
	e.tags = append(e.tags, tags...)
	return e
}

// Add appends one tag to the list of tags.
func (e *Error) Add(key string, value interface{}) *Error {
	return e.AddTags(Tags{key: value})
}

// SetCode adds a status code to the error
func (e *Error) SetCode(code Code) *Error {
	e.code = code
	return e
}

// SetMessage adds a final friendly message to the error to be returned to client
func (e *Error) SetMessage(msg string) *Error {
	e.clientMsg = msg
	return e
}

// Code returns the embedded error status code
func (e *Error) Code() Code {
	return e.code
}

// RawTags returns the raw list of tags.
// This function only needs to be used if you may have duplicates
// or want to look at the individual tag maps.
func (e *Error) RawTags() []Tags {
	return e.tags
}

// Tags returns a merged view of the tags attached to this error.
// Tags at creation override tags added later, since that is assumed
// to be the authoritative source.
func (e *Error) Tags() map[string]interface{} {
	tags := make(map[string]interface{})
	for i := len(e.tags) - 1; i >= 0; i-- {
		for k, v := range e.tags[i] {
			tags[k] = v
		}
	}
	return tags
}

// Stack returns the stack at time of creation.
func (e *Error) Stack() []StackEntry {
	entries := make([]StackEntry, len(e.stack))
	for i, pc := range e.stack {
		// see documentation for runtime.Callers for the pc-1
		// TODO: The function after runtime.sigpanic shouldn't be -1
		f := runtime.FuncForPC(pc - 1)
		if f != nil {
			file, line := f.FileLine(pc - 1)
			entries[i] = StackEntry{
				Name: f.Name(),
				File: file,
				Line: line,
			}
		}
	}
	return entries
}

func (e *Error) PrintStack() string {
	var buffer bytes.Buffer

	for _, frame := range e.Stack() {
		buffer.WriteString(frame.String())
		buffer.WriteString("\n")
	}

	return buffer.String()
}

// PrintStack is a way to call *Error.PrintStack when all you have is a standard error
func PrintStack(err error) string {
	if e, ok := err.(*Error); ok {
		return e.PrintStack()
	}

	return ""
}

// Is returns true if both errors refer to the same error.
func (e *Error) Is(err error) bool {
	return e.err == unwrapError(err)
}

// Tags are attached to the errors. This is used for better error messages in Sentry.
type Tags map[string]interface{}

// One entry in the stacktrace.
type StackEntry struct {
	// Name of the function
	Name string

	// File the function is in
	File string

	// Line number
	Line int
}

func (e StackEntry) String() string {
	return fmt.Sprintf("%s()\n\t%s:%v", e.Name, e.File, e.Line)
}

// FinalWrap is used as the last time an error gets wrapped before
// being returned to the handler, and then the client
func FinalWrap(err error, code Code, msg string) *Error {
	if werr, ok := err.(*Error); ok {
		werr.SetCode(code)
		werr.SetMessage(msg)
		return werr
	}

	return &Error{
		err:       err,
		code:      code,
		clientMsg: msg,
	}
}

// New creates a new wrapped error.
func New(msg string) *Error {
	e := &Error{
		err: errors.New(msg),
	}

	e.addStack()

	return e
}

// Is returns true if both unwrapped errors are the same.
func Is(err error, err2 error) bool {
	return unwrapError(err) == unwrapError(err2)
}

// Cast returns the Error or nil if it is not wrapped.
// This should only be used by code that wants to report the error message
func Cast(err error) *Error {
	berr, _ := err.(*Error)
	return berr
}

func (e *Error) addStack(offset ...int) {
	stack := make([]uintptr, 16)
	// skip 3: runtime.Callers, wrap(), New() or Wrap()
	n := 3
	if len(offset) > 0 {
		n += offset[0]
	}
	l := runtime.Callers(n, stack)
	e.stack = stack[:l]
}

func unwrapError(err error) error {

	for {
		if berr, ok := err.(*Error); ok {
			err = berr.err
			continue
		}
		return err
	}
}

// Cause returns the root cause of the error
func Cause(err error) error {
	if berr, ok := err.(*Error); ok {
		return berr.err
	}

	return nil
}
