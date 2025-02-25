#!/usr/bin/env bash

set +e  # Disable exit on error, we want to clean up after an error

export AWS_ACCESS_KEY_ID="test"
export AWS_SECRET_ACCESS_KEY="test"
export AWS_DEFAULT_REGION="eu-west-1"
export AWS_ENDPOINT_URL="http://localhost:4566"

uuid=$(uuidgen)

# get the queue URL
queue_url=$(aws sqs get-queue-url --queue-name result-queue | jq -r '.QueueUrl')
echo "Queue URL: $queue_url"

# Background process: posts item to SQS queue to trigger the response
(
  # Generate random number between 1-5
  sleep_time=$(( (RANDOM % 5) + 1 ))
  echo "Will execute AWS command after ${sleep_time} seconds..."
  sleep ${sleep_time}
  aws sqs send-message \
    --endpoint-url http://localhost:4566 \
    --queue-url http://localhost:4566/000000000000/result-queue \
    --message-body "{\"uuid\": \"$uuid\", \"payload\": {\"name\": \"response payload from service\"}}"
  echo "AWS command executed."
) &
background_pid=$!

echo "Request UUID: $uuid"

curl -XPOST --retry 5 --fail-with-body \
    localhost:8080/api/process \
    -H "Content-Type: application/json" \
    -d "{\"uuid\": \"$uuid\", \"payload\": {\"name\": \"default\", \"timeout\": 20}}"
curl_exit_status=$?

if [ $curl_exit_status -ne 0 ]; then
  echo "\nCurl failed with status $curl_exit_status. Killing background process..."
  kill $background_pid
  wait $background_pid 2>/dev/null || true
  echo "Background process terminated."
  echo "Test failed."
  exit $curl_exit_status
fi

echo "\nWaiting for AWS command to complete..."
wait $background_pid
echo "Test successful."
