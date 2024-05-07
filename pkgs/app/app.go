// interface for applications
package app

import "context"

type App interface {
	Starter
	Stopper
}

type Starter interface {
	Start(context.Context) error
}

type Stopper interface {
	Stop(context.Context) error
}
