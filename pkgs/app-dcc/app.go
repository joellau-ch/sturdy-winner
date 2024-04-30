package appdcc

import (
	"context"

	interfaceapp "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/interface-app"
	interfacepubsub "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/interface-pub-sub"
	simplemessagerouter "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/simple-message-router"
	"github.com/pkg/errors"
)

// ================
// At a high level, this file is only concerned that
// `DccApp` is composed of multiple sub-applications
// (MaxxTrader FIX Client, Talos WS Client)
//
// See: [Composite Pattern](https://refactoring.guru/design-patterns/composite)
//
// ================
//
//	                      --------
//	                     | DccApp |
//	                      --------
//	                         |
//	       .-------------------------------------.
//	       v                                     v
//	-----------------                -----------------------
//
// |                 |              |                       |
// | Talos WS Client |              | MaxxTrader FIX Client |
// |                 |              |                       |
//
//	-----------------                -----------------------
type DccApp struct {
	Apps          AppRegistry
	MessageRouter MessageRouter
}

var _ interfaceapp.App = &DccApp{}

type MessageRouter interfacepubsub.PubSubber[any]

type NewAppDefinition struct {
	Name string
	New  AppFactoryFunc
}

var apps = []NewAppDefinition{
	{
		"MaxxTraderFixClient",
		func(ps interfacepubsub.PubSubber[any]) (interfaceapp.App, error) {
			return NewMaxxTraderFixClient(ps)
		},
	},
	{
		"TalosWebsocketClient",
		func(ps interfacepubsub.PubSubber[any]) (interfaceapp.App, error) {
			return NewTalosWsClient(ps)
		},
	},
}

type AppFactoryFunc func(interfacepubsub.PubSubber[any]) (interfaceapp.App, error)

func New() (app *DccApp, err error) {
	app = &DccApp{}

	// Init MessageRouter
	app.MessageRouter = simplemessagerouter.NewSimpleMessageRouter()

	// Register Applications
	for _, newapp := range apps {
		subApp, err := newapp.New(app.MessageRouter)
		if err != nil {
			return nil, err
		}
		app.Apps.Register(newapp.Name, subApp)
	}

	return
}

func (d *DccApp) Start(ctx context.Context) (err error) {
	for name, app := range d.Apps {
		err = app.Start(ctx)
		if err != nil {
			return errors.Wrapf(err, "could not start app: %s", name)
		}
	}
	return nil
}

func (d *DccApp) Stop(ctx context.Context) (err error) {
	for name, app := range d.Apps {
		err = app.Stop(ctx)
		if err != nil {
			return errors.Wrapf(err, "could not stop app: %s", name)
		}
	}
	return nil
}
