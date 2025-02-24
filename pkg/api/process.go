package api

import (
	"log"

	"github.com/meero-com/hmb-proxy/pkg/aws"
	"github.com/meero-com/hmb-proxy/pkg/config"
	"github.com/meero-com/hmb-proxy/pkg/pollers"
	"github.com/meero-com/hmb-proxy/pkg/producers"
)

func process(ch chan string, p requestPayload) {
	var producer producers.Producer
	var poller pollers.Poller

	backendType := config.GetConfig("backend_type").(string)

	if backendType == "ddb" {
		ddb := aws.NewDdbCoordinator()
		producer = &producers.DdbProducer{
			Ddb: ddb,
		}
		poller = &pollers.DdbPoller{
			Ddb: ddb,
		}
	} else {
		// TODO: sqs producer and poller
		log.Fatalf("SQS backend not yet implemented")
	}

	producer.Produce(p.Uuid, p.Payload.Name)
	poller.Poll(ch, p.Uuid)
}
