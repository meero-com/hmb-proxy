#!/bin/bash
DLQ_SQS=input-queue-dlq
SOURCE_SQS=input-queue

QUEUE_URL=$(awslocal sqs --region eu-west-1 create-queue --queue-name $DLQ_SQS | grep '"QueueUrl"' | awk -F '"QueueUrl":' '{print $2}' | tr -d '"' | xargs)

DLQ_SQS_ARN=$(awslocal sqs --region eu-west-1 get-queue-attributes \
                  --attribute-name QueueArn --queue-url=$QUEUE_URL \
                  |  sed 's/"QueueArn"/\n"QueueArn"/g' | grep '"QueueArn"' | awk -F '"QueueArn":' '{print $2}' | tr -d '"' | xargs)

awslocal sqs --region eu-west-1 create-queue --queue-name $SOURCE_SQS \
     --attributes '{
                   "RedrivePolicy": "{\"deadLetterTargetArn\":\"'"$DLQ_SQS_ARN"'\",\"maxReceiveCount\":\"2\"}",
                   "VisibilityTimeout": "20"
                   }'


DLQ_SQS=result-queue-dlq
SOURCE_SQS=result-queue

QUEUE_URL=$(awslocal sqs --region eu-west-1 create-queue --queue-name $DLQ_SQS | grep '"QueueUrl"' | awk -F '"QueueUrl":' '{print $2}' | tr -d '"' | xargs)

DLQ_SQS_ARN=$(awslocal sqs --region eu-west-1 get-queue-attributes \
                  --attribute-name QueueArn --queue-url=$QUEUE_URL \
                  |  sed 's/"QueueArn"/\n"QueueArn"/g' | grep '"QueueArn"' | awk -F '"QueueArn":' '{print $2}' | tr -d '"' | xargs)

awslocal sqs --region eu-west-1 create-queue --queue-name $SOURCE_SQS \
     --attributes '{
                   "RedrivePolicy": "{\"deadLetterTargetArn\":\"'"$DLQ_SQS_ARN"'\",\"maxReceiveCount\":\"2\"}",
                   "VisibilityTimeout": "20"
                   }'


awslocal dynamodb create-table \
 --table-name request-ddb-table \
 --key-schema AttributeName=uuid,KeyType=HASH \
 --attribute-definitions AttributeName=uuid,AttributeType=S \
 --billing-mode PAY_PER_REQUEST \

awslocal dynamodb create-table \
 --table-name response-ddb-table \
 --key-schema AttributeName=uuid,KeyType=HASH \
 --attribute-definitions AttributeName=uuid,AttributeType=S \
 --billing-mode PAY_PER_REQUEST \

awslocal dynamodb put-item \
  --table-name request-ddb-table \
  --item '{"uuid": {"S": "c011892b-2204-4b39-89c0-f4f67a905cd2"}, "payload": {"M": {"name": {"S": "template test"}}}}'
