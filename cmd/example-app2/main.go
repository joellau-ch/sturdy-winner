// example of 2nd executable file within same project
package main

import (
	"github.com/coinhako/joellau-ch/sturdy-winner/pkgs/app2"
)

func main() {
	app := app2.App2{}
	app.Start()
}
