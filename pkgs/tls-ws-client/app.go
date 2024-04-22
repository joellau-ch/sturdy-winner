package tlswsclient

import (
	"context"

	interfaceapp "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/interface-app"
)

type TalosApp struct {
	OnExecutionReport       func(executionreport any) error
	OnClientExecutionReport func(executionreport any) error
	// TODO: list other hooks here
}

var _ interfaceapp.App = &TalosApp{}

func New() (app *TalosApp, err error) {
	app = &TalosApp{}

	return
}

func (app *TalosApp) Start(context.Context) error {
	return nil
}

func (app *TalosApp) Stop(context.Context) error {
	return nil
}
