#!/bin/bash
aws sqs --endpoint http://localhost:4566 send-message --queue-url http://sqs.eu-west-1.localhost.localstack.cloud:4566/000000000000/input-queue --message-body 'http://sqs.eu-west-1.localhost.localstack.cloud:4566/000000000000/result-queue-1234'
# aws sqs --endpoint http://localhost:4566 send-message --queue-url http://sqs.eu-west-1.localhost.localstack.cloud:4566/000000000000/input-queue --message-body '{"hello": "orld2313"}'
