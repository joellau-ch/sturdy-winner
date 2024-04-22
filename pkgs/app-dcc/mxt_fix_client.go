package appdcc

import (
	appmxtfixclient "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/app-mxt-fix-client"
	interfacepubsub "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/interface-pub-sub"
	"github.com/quickfixgo/fix44/executionreport"
)

func NewMaxxTraderFixClient(pubsub interfacepubsub.PubSubber[any]) (app *appmxtfixclient.MaxxTraderApp, err error) {
	// New MaxxTrader FIX Client
	mxtapp, err := appmxtfixclient.New()
	if err != nil {
		return
	}

	// Register Publishers
	mxtapp.ExecutionReportHandler = func(executionreport executionreport.ExecutionReport) (err error) {
		pubsub.Publish(RouteMaxxTraderExecutionReport, executionreport)
		return
	}

	return
}
