package appdcc

import (
	"fmt"

	interfaceapp "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/interface-app"
)

type AppRegistry map[string]interfaceapp.App

func (ar *AppRegistry) Register(name string, app interfaceapp.App) error {
	_, found := (*ar)[name]
	if found {
		return fmt.Errorf("name has already been taken: %s", name)
	}

	(*ar)[name] = app
	return nil
}
