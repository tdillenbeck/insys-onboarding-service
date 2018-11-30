package werror

import (
	"reflect"
)

type wrapper func(werr *Error, rawerr error) *Error

var wrappers = make(map[string]wrapper)

// register default wrapper
func init() {
	RegisterWrapper(reflect.TypeOf(&Error{}), defaultWrap)
}

/*
	RegisterWrapper registers a function that knows
	how to extract information out of an error of type t
	and convert it into a werror.Error with all of the
	meta data attached from error of type t.
*/
func RegisterWrapper(t reflect.Type, w wrapper) {
	wrappers[t.String()] = w
}

func RegisterWrapperString(s string, w wrapper) {
	wrappers[s] = w
}

/*
	Wrap appends messages and meta data from err
	onto a werror.Error
*/
func Wrap(err error, msgs ...string) *Error {

	if err == nil {
		err = New("werror.Wrap called with nil error")
	}

	t := reflect.TypeOf(err)

	var werr *Error
	w, ok := wrappers[t.String()]
	if !ok {
		werr = &Error{
			err: err,
		}

	} else {
		werr = defaultWrap(werr, err)
		werr = w(werr, err)
	}

	werr.messages = append(werr.messages, msgs...)

	if len(werr.stack) == 0 {
		werr.addStack()
	}

	return werr

}

// convert from error to Error
func defaultWrap(werr *Error, err error) *Error {

	werr, ok := err.(*Error)
	if !ok {
		return &Error{
			err:      err,
			messages: []string{},
		}
	}

	return werr
}
