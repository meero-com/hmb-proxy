package producers

import (
	"context"
	"log"

	"github.com/meero-com/hmb-proxy/pkg/aws"
	"github.com/meero-com/hmb-proxy/pkg/config"
)

type SqsProducer struct {
	Sqs aws.SqsCoordinator
}

func (p *SqsProducer) Produce(uuid string, payload string) {
	sqs := p.Sqs
	queueUrl := config.GetConfig("sqs.source_queue").(string)

	// TODO: forge sqs payload with uuid in message attributes
	sqsPayload := payload

	err := sqs.PutMessage(context.Background(), queueUrl, &sqsPayload)

	if err != nil {
		log.Fatalf("Failed to put message %s to sqs queue %s", uuid, queueUrl)
	}
}
