# HMB proxy

<p align="center">
  <img src="./logo.png" alt="logo.png" width="250" height="250"><br>
  <b>Hold My Beer proxy</b>
</p>

------

Abstract asynchronous processing and make it synchronous for your clients.

[![CI](https://github.com/meero-com/hmb-proxy/actions/workflows/ci.yml/badge.svg)](https://github.com/meero-com/hmb-proxy/actions/workflows/ci.yml)

## Goal

The HMB proxy is made to allow under-the-hood, loosely coupled processing with parallel processing in mind.

Instead of having a "fire & forget", callback-supported asynchronous processing, customers can rely on simple HTTP calls
while having all the benefits of async processing.

Using a **timeout** / **retry** / **exponential back-off** strategy, users can completely abstract the implementation
complexity of asynchronous systems.

<p align="center">
  <img src="./docs/images/hmb-full.png" alt="logo.png"><br>
</p>

## Usage

```console
$ export GIN_MODE=release
$ hmb-proxy --env=prod
2025/02/19 15:33:01 sqs: map[destination_queue:output-queue source_queue:input-queue]
2025/02/19 15:33:01 aws: map[access_key_id:default endpoint_url:http://localstack:4566/ region:eu-west-1 secret_access_key:default]
2025/02/19 15:33:01 env: prod
2025/02/19 15:33:01 server: map[port:8080]
```

You can then forward requests to the proxy using `curl`.

A dummy-service is available in [./dummy-service](./dummy-service) to try out the
loosely coupling mechanism locally.

Alternatively, the repository contains a [`./e2e`](./e2e) script to manually simulate
a backend service.

> This project is not Production-ready as of now. Feel free to use it, give your feedback and to help us improve it !

## Build

This project supports packaging through container images.

The proxy can be built using the following command:

```console
$ docker build -t hmb-proxy:local .
```

It can then be run locally using:
```console
$ docker run -it --rm hmb-proxy:local
```

## Development

The repository offers different utilities to improve development.

### Compose setup

A `docker-compose.yml` manifest is available in the top-level directory to emulate AWS services using localstack.

```console
$ docker compose up -d --wait
[+] Running 2/2
 ✔ Container hmb-proxy-hmb-proxy-1   Healthy  1.8s
 ✔ Container hmb-proxy-localstack-1  Healthy  6.8s
```

## Glossary

- Project name: HMB proxy
- Reference to the executable: hmb-proxy
