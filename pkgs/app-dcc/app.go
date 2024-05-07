package appdcc

import (
	"context"

	interfaceapp "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/app"
	msgrtr "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/message-router"
	"github.com/pkg/errors"
)

// ================
// At a high level, this file is only concerned that
// `DccApp` is composed of multiple sub-applications
// (MaxxTrader FIX Client, Talos WS Client)
//
// See: [Composite Pattern](https://refactoring.guru/design-patterns/composite)
type DccApp struct {
	Apps          AppRegistry
	MessageRouter msgrtr.MessageRouter[RouteKey, RouteMessage, RouteSubId]
}

var _ interfaceapp.App = &DccApp{}

type RouteKey = msgrtr.RouteKey
type RouteMessage = msgrtr.Message
type RouteSubId = msgrtr.SubscriptionId

func New() (app *DccApp, err error) {
	app = &DccApp{}
	router, err := msgrtr.NewSimpleMessageRouter()
	app.MessageRouter = router

	return
}

func (d *DccApp) Start(ctx context.Context) (err error) {
	for name, app := range d.Apps {
		err = app.Start(ctx)
		if err != nil {
			return errors.Wrapf(err, "could not start app: %s", name)
		}
	}

	return
}

func (d *DccApp) Stop(ctx context.Context) (err error) {
	for name, app := range d.Apps {
		err = app.Stop(ctx)
		if err != nil {
			return errors.Wrapf(err, "could not stop app: %s", name)
		}
	}

	return
}
