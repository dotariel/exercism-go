package erratum

import (
	"io"
)

const testVersion = 2

type TransientError struct {
	err error
}

func (e TransientError) Error() string {
	return e.err.Error()
}

type FrobError struct {
	defrobTag string
	inner     error
}

func (e FrobError) Error() string {
	return e.inner.Error()
}

type Resource interface {
	io.Closer

	// Frob does something with the input string.
	// Because this is an incredibly badly designed system if there is an error
	// it panics.
	//
	// The paniced error may be a FrobError in which case Defrob should be called
	// with the defrobTag string.
	Frob(string)

	Defrob(string)
}

// ResourceOpener is a function that creates a resource.
//
// It may return a wrapped error of type TransientError. In this case the resource
// is temporarily unavailable and the caller should retry soon.
type ResourceOpener func() (Resource, error)

func Use(o ResourceOpener, input string) (err error) {

	res, err := o()

	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return err
	}

	defer func() {
		defer res.Close()
		if p := recover(); p != nil {
			if ferr, ok := p.(FrobError); ok {
				res.Defrob(ferr.defrobTag)
			}
			err = p.(error)
		}
	}()

	res.Frob(input)
	return err
}
