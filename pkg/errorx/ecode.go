package errorx

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type Codes interface {
	// Code get error code.
	Code() int
	// Message get code message.
	Message() string
}

var (
	_messages = map[int]string{}
	_codes    = map[int]struct{}{} // register codes.
)

// A Code is an int error code spec.
type Code int

// New new a ecode.Codes by int value.
// NOTE: ecode must unique in global, the New will check repeat and then panic.
func New(e int, msg string) Code {
	if e < 0 {
		panic("ecode must be greater or equal than zero")
	}
	return add(e, msg)
}

func add(e int, msg string) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	_messages[e] = msg
	return Int(e)
}

func (e Code) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

// Code return error code
func (e Code) Code() int { return int(e) }

// Message return error message
func (e Code) Message() string {
	if msg, ok := _messages[e.Code()]; ok {
		return msg
	}
	return e.Error()
}

// Cause cause from error to ecode.
func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return String(e.Error())
}

// Int parse code int to error.
func Int(i int) Code { return Code(i) }

// String parse code string to error.
func String(e string) Code {
	if len(e) == 0 {
		return OK
	}
	// try error string
	i, err := strconv.Atoi(e)
	if err != nil {
		return ServerErr
	}
	return Code(i)
}

// Equal equal a and b by code int.
func Equal(a, b Codes) bool {
	if a == nil {
		a = OK
	}
	if b == nil {
		b = OK
	}
	return a.Code() == b.Code()
}

// EqualError equal error
func EqualError(code Codes, err error) bool {
	return Cause(err).Code() == code.Code()
}
