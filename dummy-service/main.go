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

type Payload struct {
}

type SqsCoordinator struct {
	SqsClient *sqs.Client
}

func (s SqsCoordinator) GetMessages(ctx context.Context, queueUrl string, maxMessages int32, waitTime int32) ([]types.Message, error) {
	var messages []types.Message
	result, err := s.SqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
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

func (s SqsCoordinator) AckMessage(ctx context.Context, queueUrl string, receiptHandle *string) error {
	_, err := s.SqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueUrl),
		ReceiptHandle: receiptHandle,
	})
	if err != nil {
		log.Printf("Couldn't acknowledge messages from queue=%v: %v\n", queueUrl, err)
	} else {
		log.Printf("pew pew the message is pew'd")
	}
	return err
}

func (s SqsCoordinator) PutMessage(ctx context.Context, queueUrl string, body *string) error {
	response, err := s.SqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueUrl),
		MessageBody: body,
	})
	if err != nil {
		log.Printf("Couldn't get messages from queue %v. Here's why: %v\n", queueUrl, err)
	} else {
		log.Printf("sent the message to queue %s %v", queueUrl, *response.MessageId)
	}
	return err
}

func main() {
	inputQueue := os.Getenv("SQS_SOURCE_QUEUE")
	// responseQueue := os.Getenv("SQS_RESPONSE_QUEUE") // TODO: rely on input message
	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("Could not load config: %v\n", err)
		return
	}
	sqsClient := sqs.NewFromConfig(sdkConfig)
	action := SqsCoordinator{SqsClient: sqsClient}
	log.Printf("start polling\n")
	for {
		messages, _ := action.GetMessages(context.Background(), inputQueue, 10, 10)
		log.Printf("got %d messages \n", len(messages))
		for idx, message := range messages {
			log.Printf("got messages #%d id=%s %s\n", idx, *message.MessageId, *message.Body)

			// forge response payload

			// send response
			dummyResponse := "random answer"
			// responseQueue
			action.PutMessage(context.Background(), *message.Body, &dummyResponse)

			action.AckMessage(context.Background(), inputQueue, message.ReceiptHandle)
		}
	}
}
