package integrationtest

import (
	"github.com/state-alchemists/ayanami-service-go/config"
	"github.com/state-alchemists/ayanami-service-go/gateway"
	"github.com/state-alchemists/ayanami-service-go/msgbroker"
	"log"
)

// MainGateway emulating gateway's main
func MainGateway() {
	routes := []string{ // define your routes here
		"/",
		"/banner/",
	}
	broker, err := msgbroker.NewNats(config.GetNatsURL())
	if err != nil {
		log.Fatal(err)
	}
	port := config.GetGatewayPort()
	multipartFormLimit := config.GetGatewayMultipartFormLimit()
	gateway.Serve(broker, port, multipartFormLimit, routes)
}
