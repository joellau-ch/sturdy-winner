// main executable
package main

import (
	"context"
	"log"

	appdcc "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/app-dcc"
)

func main() {
	ctx := context.Background()

	log.Println("initializing application")
	app, err := appdcc.New()
	if err != nil {
		log.Fatalf("could not create new dcc app: %v", err)
	}

	log.Println("starting application")
	err = app.Start(ctx)
	if err != nil {
		log.Fatalf("could not start dcc app: %v", err)
	}
	defer func(app *appdcc.DccApp, ctx context.Context) {
		log.Println("attempting graceful shutdown")
		err := app.Stop(ctx)
		if err != nil {
			log.Fatalf("error stopping DccApp: %v", err)
		}
	}(app, ctx)
}
