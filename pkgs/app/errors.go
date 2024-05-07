package app

import "fmt"

type ErrCouldNotStart struct {
	AppName string
}

var _ error = &ErrCouldNotStart{}

func (e *ErrCouldNotStart) Error() string {
	return fmt.Sprintf("Could not start application: %+v\n", e.AppName)
}
