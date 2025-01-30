#!/usr/bin/env bash

curl -v -H "Content-Type: application/json" -XPOST localhost:8080/api/test -d '{"uuid": "c011892b-2204-4b39-89c0-f4f67a905cd2", "payload": {"name": "default", "timeout": 2}}'
