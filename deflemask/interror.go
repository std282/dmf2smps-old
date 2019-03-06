package deflemask

import "fmt"

type internalError struct {
	underlyingError error
}

func (prerr internalError) Error() string {
	return prerr.underlyingError.Error()
}

func internalPanic(err error) {
	panic(internalError{underlyingError: err})
}

func internalPanicNewf(format string, args ...interface{}) {
	internalPanic(fmt.Errorf(format, args...))
}

func internalRecover() error {
	e := recover()

	if e == nil {
		return nil
	}

	if err, ok := e.(internalError); ok {
		return err.underlyingError
	}

	panic(e)
}
