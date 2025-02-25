#!/usr/bin/env bash

set +e  # Disable exit on error, we want to clean up after an error

export AWS_ACCESS_KEY_ID="test"
export AWS_SECRET_ACCESS_KEY="test"
export AWS_DEFAULT_REGION="eu-west-1"
export AWS_ENDPOINT_URL="http://localhost:4566"

uuid=$(uuidgen)

# Background process: posts item to DynamoDB to trigger the response
(
  # Generate random number between 1-5
  sleep_time=$(( (RANDOM % 5) + 1 ))
  echo "Will execute AWS command after ${sleep_time} seconds..."
  sleep ${sleep_time}
  aws dynamodb put-item \
    --endpoint-url http://localhost:4566 \
    --table-name response-ddb-table \
    --item '{"uuid": {"S": "'$uuid'"}, "payload": { "M": { "name": { "S": "response payload from service" } } }}'
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
