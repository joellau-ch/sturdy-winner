package dcapp

import "fmt"

type DccApp struct{}

func (d *DccApp) Start() {
	fmt.Println("dccapp")
}
