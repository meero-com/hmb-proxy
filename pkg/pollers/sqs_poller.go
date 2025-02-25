package pollers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/meero-com/hmb-proxy/pkg/aws"
	"github.com/meero-com/hmb-proxy/pkg/config"
)

type SqsPoller struct {
	Sqs aws.SqsCoordinator
}

func (p *SqsPoller) Poll(ch chan string, uuid string) {
	c := p.Sqs
	destinationQueue := config.GetConfig("sqs.destination_queue").(string)

	log.Printf("Start polling\n")

	for {
		messages, _ := c.GetMessages(context.Background(), destinationQueue, 10, 10)
		log.Printf("got %d messages \n", len(messages))
		for idx, message := range messages {
			log.Printf("got messages #%d id=%s %s\n", idx, *message.MessageId, *message.Body)

			// TODO: check if message attributes contain uuid, if not reschedule

			err := c.AckMessage(context.Background(), destinationQueue, message.ReceiptHandle)
			if err != nil {
				log.Fatalf("Failed to acknowledge message %s", *message.MessageId)
			}

			si, err := json.Marshal(*message.Body)
			if err != nil {
				log.Fatalf("Failed to Marshal item: %s", *message.Body)
			}
			ch <- string(si)
			close(ch)
			return
		}

		time.Sleep(pollInterval)
	}
}
