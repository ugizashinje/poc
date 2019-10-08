package failures

import (
	"fmt"
	"runtime"
)

type Error struct {
	err    error
	notify bool
	Text   string   `json:"error"`
	Stack  []string `json:"-"`
	status int
}

func create() *Error {
	pcs := make([]uintptr, 64)
	height := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs)
	stack := []string{}
	for frame, more := frames.Next(); more && height > 0; frame, more = frames.Next() {
		height--
		stack = append(stack, fmt.Sprintf("%s %s:%d", frame.Function, frame.File, frame.Line))
	}

	fmt.Println(stack)
	return &Error{Stack: stack,
		notify: false,
	}
}
func (e *Error) Error() string {
	return e.Text
}
func (e *Error) Status() int {
	return e.status
}
func (e *Error) WithInfo(format string, args ...interface{}) SuperError {
	reFormat := "%s " + format
	reArgs := []interface{}{e.Text, args}
	e.Text = fmt.Sprintf(reFormat, reArgs)
	return e
}

func (e *Error) Notify() SuperError {
	e.notify = true
	return e
}

func (e *Error) String() string {
	return e.Text
}

type SuperError interface {
	Error() string
	WithInfo(fmt string, args ...interface{}) SuperError
	Notify() SuperError
	String() string
	Status() int
}
