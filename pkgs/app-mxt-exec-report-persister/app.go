package mxter

import (
	"context"

	appi "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/app-utils"
	smr "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/message-router"
)

// ================
// [M]ax[xt]rader [Exec]ution [R]eport [P]ersister [App]
//
// persists execution reports from Maxtrader
type MxtExecRptPApp struct {
	subscriptions map[SubKey]bool // set of subscription ids

	MessageRouter smr.MessageRouter[smr.RouteKey, smr.Message, smr.SubscriptionId]
}

type SubKey struct {
	smr.RouteKey
	smr.SubscriptionId
}

var _ appi.App = &MxtExecRptPApp{}

func NewMxtExecRptPApp(router smr.MessageRouter[smr.RouteKey, smr.Message, smr.SubscriptionId]) (app *MxtExecRptPApp, err error) {
	app = &MxtExecRptPApp{}
	app.subscriptions = map[SubKey]bool{}

	app.MessageRouter = router

	return
}

func (app *MxtExecRptPApp) Start(ctx context.Context) error {
	// subscribe handler to topic
	key := smr.RouteKey{Route: "/maxxtrader/execution-reports", Type: "ExecutionReport"}
	subid, err := app.MessageRouter.Subscribe(key, app.OnExecutionReport)
	if err != nil {
		err = app.Stop(ctx)
		return err
	}
	app.subscriptions[SubKey{key, subid}] = true

	return nil
}

func (app *MxtExecRptPApp) Stop(context.Context) error {
	// unsubscribe from all topics
	for subkey := range app.subscriptions {
		delete(app.subscriptions, subkey)
	}

	return nil
}

// transform and persist message
func (app *MxtExecRptPApp) OnExecutionReport(message smr.Message) error {
	// TODO: immplement
	return nil
}
