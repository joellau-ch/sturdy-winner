// TODO: implement ws interfaces for connecting with talos
package appdcc

import (
	interfacepubsub "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/interface-pub-sub"
	tlswsclient "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/tls-ws-client"
)

func NewTalosWsClient(pubsub interfacepubsub.PubSubber[any]) (app *tlswsclient.TalosApp, err error) {
	// New Talos WS App
	talosApp, err := tlswsclient.New()
	if err != nil {
		return
	}

	// Register Publishers
	talosApp.OnExecutionReport = func(executionreport any) error {
		pubsub.Publish("asdf", executionreport)
		return nil
	}
	talosApp.OnClientExecutionReport = func(clientexecutionreport any) error {
		pubsub.Publish(RouteTalosClientExecutionReport, clientexecutionreport)
		return nil
	}

	return
}
