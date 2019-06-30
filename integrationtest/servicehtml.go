package integrationtest

import (
	"github.com/state-alchemists/ayanami-service-go/config"
	"github.com/state-alchemists/ayanami-service-go/msgbroker"
	"github.com/state-alchemists/ayanami-service-go/service"
	"log"
)

// MainServiceHTML emulating HTML's main
func MainServiceHTML() {
	serviceName := "html"
	// define broker
	broker, err := msgbroker.NewNats(config.GetNatsURL())
	if err != nil {
		log.Fatal(err)
	}
	// define services
	services := service.Services{
		service.NewService(serviceName, "pre",
			[]string{"input"},
			[]string{"output"},
			WrappedPre,
		),
	}
	// consume and publish forever
	forever := make(chan bool)
	services.ConsumeAndPublish(broker, serviceName)
	<-forever
}
