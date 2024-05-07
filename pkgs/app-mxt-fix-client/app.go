// TODO: implement quickfix interfaces for connecting with maxxtrader
package appmxtfixclient

import (
	"context"

	"github.com/coinhako/joellau-ch/sturdy-winner/pkgs/app"
	simplemessagerouter "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/message-router"
)

type MaxxTraderApp struct {
	// TODO: list other hooks here
	MessageRouter *simplemessagerouter.MessageRouter[simplemessagerouter.RouteKey, simplemessagerouter.Message, simplemessagerouter.SubscriptionId]
}

var _ app.App = &MaxxTraderApp{}

func New() (app *MaxxTraderApp, err error) {
	app = &MaxxTraderApp{}

	return
}

func (app *MaxxTraderApp) Start(context.Context) error {
	return nil
}

func (app *MaxxTraderApp) Stop(context.Context) error {
	return nil
}
