package apphub

import (
	"fmt"

	appi "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/app-utils"
)

type AppRegistry map[string]appi.App

func (ar *AppRegistry) Register(name string, app appi.App) error {
	_, found := (*ar)[name]
	if found {
		return fmt.Errorf("name has already been taken: %s", name)
	}

	(*ar)[name] = app
	return nil
}
