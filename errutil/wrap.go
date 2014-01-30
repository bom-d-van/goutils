package errutil

import (
	"fmt"
	"runtime"
	"time"
)

type Err struct {
	err     string
	details string
}

func (e Err) Error() string {
	return e.err
}

func (e Err) Details() string {
	return e.details
}

// for test
var getNow = func() string { return time.Now().Format("2006/01/02 15:04:05") }

// Always return an Err.
func Wrap(err interface{}) error {
	_, file, line, _ := runtime.Caller(1)
	if errs, ok := err.(Err); ok {
		errs.details += fmt.Sprintf("\n%s:%d", file, line)
		err = errs
	} else {
		if _, ok := err.(error); !ok {
			err = fmt.Errorf("%+v", err)
		}
		marker := fmt.Sprintf("[Error Tracks %s]", getNow())
		msg := err.(error).Error()
		err = Err{
			err:     msg,
			details: fmt.Sprintf("%s %s\n%s:%d", marker, msg, file, line),
		}
	}

	return err.(error)
}
