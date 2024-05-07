package tlswsclient

import (
	"context"

	appi "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/app-utils"
	smr "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/message-router"
)

type TalosDropCopyApp struct {
	MessageRouter smr.MessageRouter[smr.RouteKey, smr.Message, smr.SubscriptionId]
}

var _ appi.App = &TalosDropCopyApp{}

func NewTalosApp(router smr.MessageRouter[smr.RouteKey, smr.Message, smr.SubscriptionId]) (app *TalosDropCopyApp, err error) {
	app = &TalosDropCopyApp{}

	app.MessageRouter = router

	return
}

func (app *TalosDropCopyApp) Start(context.Context) error {
	return nil
}

func (app *TalosDropCopyApp) Stop(context.Context) error {
	return nil
}
