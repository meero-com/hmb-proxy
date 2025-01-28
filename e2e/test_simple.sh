#!/usr/bin/env bash

curl -v -H "Content-Type: application/json" -XPOST localhost:8080/api/test -d '{"id": "random-id", "name": "default"}'
