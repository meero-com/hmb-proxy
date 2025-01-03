package pollers

import (
	"context"
	"log"
	"os"

	"github.com/meero-com/guild-proxy/utils/aws"
)

func Poll() {
	inputQueue := os.Getenv("SQS_SOURCE_QUEUE")
	ctx := context.Background()
	c := aws.New(ctx)

	log.Printf("Start polling\n")

	for {
		messages, _ := c.GetMessages(context.Background(), inputQueue, 10, 10)
		log.Printf("got %d messages \n", len(messages))
		for idx, message := range messages {
			log.Printf("got messages #%d id=%s %s\n", idx, *message.MessageId, *message.Body)

			// forge response payload

			// send response
			dummyResponse := "random answer"
			// responseQueue
			c.PutMessage(context.Background(), *message.Body, &dummyResponse)

			c.AckMessage(context.Background(), inputQueue, message.ReceiptHandle)
		}
	}
}
