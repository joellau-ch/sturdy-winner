// TODO: implement quickfix interfaces for connecting with maxxtrader
package appmxtfixclient

import (
	"context"

	interfaceapp "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/interface-app"
	"github.com/quickfixgo/fix44/executionreport"
)

type MaxxTraderApp struct {
	ExecutionReportHandler func(executionreport executionreport.ExecutionReport) error
	// TODO: list other hooks here
}

var _ interfaceapp.App = &MaxxTraderApp{}

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
