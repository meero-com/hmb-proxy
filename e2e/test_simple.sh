#!/usr/bin/env bash

uuid=$(uuidgen)
echo "Run the following command to trigger the response"
echo aws dynamodb put-item --endpoint-url http://localhost:4566 --table-name response-ddb-table  --item \'{\"uuid\": {\"S\": \"$uuid\"}, \"payload\": { \"M\": { \"name\": { \"S\": \"response payload from service\" } } }}\'
curl --retry 5 -v -H "Content-Type: application/json" -XPOST localhost:8080/api/process -d "{\"uuid\": \"$uuid\", \"payload\": {\"name\": \"default\", \"timeout\": 20}}"
