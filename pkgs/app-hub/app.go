package apphub

import (
	"context"

	appi "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/app-utils"
	msgrtr "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/message-router"
	"github.com/pkg/errors"
)

// ================
// At a high level, `HubApp` is composed of multiple sub-applications
// and does not have any specific business logic of its own.
//
// It starts and stops all sub-applications in parallel.
//
// See: [Composite Pattern](https://refactoring.guru/design-patterns/composite)
type HubApp struct {
	Apps          AppRegistry
	MessageRouter msgrtr.MessageRouter[RouteKey, RouteMessage, RouteSubId]
}

var _ appi.App = &HubApp{}

type RouteKey = msgrtr.RouteKey
type RouteMessage = msgrtr.Message
type RouteSubId = msgrtr.SubscriptionId

func NewHubApp() (hub *HubApp, err error) {
	hub = &HubApp{}

	hub.MessageRouter, err = msgrtr.NewSimpleMessageRouter()
	if err != nil {
		return
	}

	return
}

func (d *HubApp) Start(ctx context.Context) (err error) {
	for name, app := range d.Apps {
		err = app.Start(ctx)
		if err != nil {
			return &appi.ErrCouldNotStart{AppName: name, Err: err}
		}
	}

	return
}

func (d *HubApp) Stop(ctx context.Context) (err error) {
	for name, app := range d.Apps {
		err = app.Stop(ctx)
		if err != nil {
			return errors.Wrapf(err, "could not stop app: %s", name)
		}
	}

	return
}
