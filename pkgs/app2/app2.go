package app2

import "fmt"

type App2 struct{}

func (a *App2) Start() {
	fmt.Println("app2")
}
