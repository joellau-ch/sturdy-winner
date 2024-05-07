package appi

import "fmt"

type ErrCouldNotStart struct {
	AppName string

	Err error
}

var _ error = &ErrCouldNotStart{}

func (e *ErrCouldNotStart) Error() string {
	return fmt.Sprintf("Could not start application: %+v : %v\n", e.AppName, e.Err)
}

type ErrCouldNotStop struct {
	AppName string

	Err error
}

var _ error = &ErrCouldNotStop{}

func (e *ErrCouldNotStop) Error() string {
	return fmt.Sprintf("Could not stop application: %+v : %v\n", e.AppName, e.Err)
}
