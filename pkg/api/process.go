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
	} else if backendType == "sqs" {
		sqs := aws.NewSqsCoordinator()
		producer = &producers.SqsProducer{
			Sqs: sqs,
		}
		poller = &pollers.SqsPoller{
			Sqs: sqs,
		}
	} else {
		log.Fatalf("Unsupported backend type: %s", backendType)
	}

	producer.Produce(p.Uuid, p.Payload.Name)
	poller.Poll(ch, p.Uuid)
}
