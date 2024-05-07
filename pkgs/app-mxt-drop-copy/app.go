// TODO: implement quickfix interfaces for connecting with maxxtrader
package appmxtfixclient

import (
	"context"

	appi "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/app-utils"
	smr "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/message-router"
)

type MxtDropCopyApplication struct {
	MessageRouter *smr.MessageRouter[smr.RouteKey, smr.Message, smr.SubscriptionId]
}

var _ appi.App = &MxtDropCopyApplication{}

func NewMaxxTraderApp(router *smr.MessageRouter[smr.RouteKey, smr.Message, smr.SubscriptionId]) (app *MxtDropCopyApplication, err error) {
	app = &MxtDropCopyApplication{}

	app.MessageRouter = router

	return
}

func (app *MxtDropCopyApplication) Start(context.Context) error {
	return nil
}

func (app *MxtDropCopyApplication) Stop(context.Context) error {
	return nil
}
