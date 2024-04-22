package appdcc

import (
	"context"

	interfaceapp "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/interface-app"
	interfacepubsub "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/interface-pub-sub"
	simplemessagerouter "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/simple-message-router"
)

// ================
// `DccApp` is an `application` that contains other applications
// (Mediator / Coordinator pattern)
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
//
// ================
// Additional Notes
// ================
//
//   - Each of these sub-applications act as *Publishers*
//
//   - We attach *Subscribers* to each topic / route that will
//     carry out business use cases (transform & persistance)
type DccApp struct {
	maxxTraderApp interfaceapp.App
	talosApp      interfaceapp.App

	MessageRouter interfacepubsub.PubSubber[any]
}

var _ interfaceapp.App = &DccApp{}

func New() (app *DccApp, err error) {
	app = &DccApp{}

	// Init MessageRouter
	messageRouter := simplemessagerouter.NewSimpleMessageRouter()

	// Setup MaxxTrader FIX Client
	go func(messageRouter interfacepubsub.PubSubber[any]) {
		if app.maxxTraderApp, err = NewMaxxTraderFixClient(messageRouter); err != nil {
			// TODO: handle error
			return
		}
	}(messageRouter)

	// Setup Talos WS Client
	go func(messageRouter interfacepubsub.PubSubber[any]) {
		if app.talosApp, err = NewTalosWsClient(messageRouter); err != nil {
			// TODO: handle error
			return
		}
	}(messageRouter)

	// Register Routes (TODO: find where to put these, avoid `any` where possible)
	messageRouter.Subscribe(RouteMaxxTraderExecutionReport, HandleMaxxTraderExecutionReport)
	messageRouter.Subscribe(RouteTalosExecutionReport, HandleTalosExecutionReport)
	messageRouter.Subscribe(RouteTalosClientExecutionReport, HandleTalosClientExecutionReport)

	// Setup MessageRouter
	app.MessageRouter = messageRouter

	return
}

func (d *DccApp) Start(ctx context.Context) (err error) {
	if err = d.maxxTraderApp.Start(ctx); err != nil {
		return
	}

	if err = d.talosApp.Start(ctx); err != nil {
		return
	}
	return nil
}

func (d *DccApp) Stop(ctx context.Context) (err error) {
	if err = d.maxxTraderApp.Stop(ctx); err != nil {
		return
	}

	if err = d.talosApp.Stop(ctx); err != nil {
		return
	}

	return nil
}
