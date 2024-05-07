// erorr types and helpers
package smr

import (
	"fmt"
)

type ErrTypeMismatch struct {
	Expected string
	Received string
}

var _ error = ErrTypeMismatch{}

func (e ErrTypeMismatch) Error() string {
	return fmt.Sprintf("Type mismatch error!: expected %+v, received %+v", e.Expected, e.Received)
}

type ErrRouteKeyCollision struct{ Key RouteKey }

var _ error = ErrRouteKeyCollision{}

func (e ErrRouteKeyCollision) Error() string {
	return fmt.Sprintf("Key already exists!: %s", e.Key)
}
