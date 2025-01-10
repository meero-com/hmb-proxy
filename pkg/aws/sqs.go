package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SqsCoordinator struct {
	SqsClient *sqs.Client
}

func NewSqsCoordinator() SqsCoordinator {
	ctx := context.Background()
	sdkCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("Could not load config: %s\n", err.Error())
	}
	sqsClient := sqs.NewFromConfig(sdkCfg)
	return SqsCoordinator{
		SqsClient: sqsClient,
	}
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
