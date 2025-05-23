package errors

import "fmt"

type PressError struct {
	internalMessage string
	publicMessage   string

	code string
}

var defaultErrorMessage = "Something went wrong"
var defaultCode = "internal_error"

func New(message string) *PressError {
	return &PressError{
		internalMessage: message,
		publicMessage:   defaultErrorMessage,
		code:            defaultCode,
	}
}

func Newf(format string, params ...interface{}) Interface {
	return New(fmt.Sprintf(format, params...))
}

func Internal(message string) Interface {
	return New(message)
}

func Internalf(format string, params ...interface{}) Interface {
	return Newf(format, params...)
}

func Public(message string) Interface {
	return &PressError{
		internalMessage: defaultErrorMessage,
		publicMessage:   message,
		code:            defaultCode,
	}
}

func Publicf(format string, params ...interface{}) Interface {
	return Public(fmt.Sprintf(format, params...))
}

func (e *PressError) Error() string {
	return e.internalMessage
}

func (e *PressError) Internal() string {
	return e.internalMessage
}

func (e *PressError) Public() string {
	return e.publicMessage
}

func (e *PressError) WithPublic(message string) Interface {
	e.publicMessage = message

	return e
}
func (e *PressError) WithInternal(message string) Interface {
	e.internalMessage = message

	return e
}
func (e *PressError) WithCode(code string) Interface {
	e.code = code

	return e
}
