// main executable
package main

import (
	"context"
	"log"

	apphub "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/app-hub"
)

func main() {
	ctx := context.Background()

	log.Println("initializing application")
	app, err := apphub.NewHubApp()
	if err != nil {
		log.Fatalf("could not create new dcc app: %v", err)
	}

	log.Println("starting application")
	err = app.Start(ctx)
	if err != nil {
		log.Fatalf("could not start dcc app: %v", err)
	}
	defer func(app *apphub.HubApp, ctx context.Context) {
		log.Println("attempting graceful shutdown")
		err := app.Stop(ctx)
		if err != nil {
			log.Fatalf("error stopping DccApp: %v", err)
		}
	}(app, ctx)
}
