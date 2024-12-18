package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/aws/aws-sdk-go/aws"
)

type SqsActions struct {
	SqsClient *sqs.Client
}

// GetMessages uses the ReceiveMessage action to get messages from an Amazon SQS queue.
func (actor SqsActions) GetMessages(ctx context.Context, queueUrl string, maxMessages int32, waitTime int32) ([]types.Message, error) {
	var messages []types.Message
	result, err := actor.SqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: maxMessages,
		WaitTimeSeconds:     waitTime,
	})
	if err != nil {
		log.Printf("Couldn't get messages from queue %v. Here's why: %v\n", queueUrl, err)
	} else {
		messages = result.Messages
	}
	return messages, err
}

func (actor SqsActions) AckMessage(ctx context.Context, queueUrl string, receiptHandle *string) error {
	_, err := actor.SqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueUrl),
		ReceiptHandle: receiptHandle,
	})
	// log.Printf("Got result %v\n", result, err)
	if err != nil {
		log.Printf("Couldn't get messages from queue %v. Here's why: %v\n", queueUrl, err)
	} else {
		log.Printf("pew pew the message is pew'd")
	}
	return err
}

func main() {
	inputQueue := os.Getenv("SQS_SOURCE_QUEUE")
	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("Could not load config: %v\n", err)
		return
	}
	sqsClient := sqs.NewFromConfig(sdkConfig)
	action := SqsActions{SqsClient: sqsClient}
	log.Printf("start polling\n")
	for {
		messages, _ := action.GetMessages(context.Background(), inputQueue, 10, 10)
		log.Printf("got messages %d\n", len(messages))
		for idx, message := range messages {
			log.Printf("got messages #%d id=%s %s\n", idx, *message.MessageId, *message.Body)
			action.AckMessage(context.Background(), inputQueue, message.ReceiptHandle)
		}
	}
}
