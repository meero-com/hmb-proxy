# Guild Proxy

Abstract asynchronous processing and make it synchronous for your clients.

## Goal

The proxy is made to allow under the hood, loosely coupled processing with parallel processing in mind.

Instead of having a "fire & forget", callback supported asynchronous processing, customers can rely on simple HTTP calls while having all the benefits of async processing.

Using a timeout / retry / exponential back-off strategy, users can completly abstract the implementation complexity of asynchronous systems.

## Build

This project supports packaging through container images.

The proxy can be built using the following command:

```console
$ docker build -t proxy:local .
```

It can then be run locally using:
```console
$ docker run -it --rm proxy:local
```

## Configure

The configuration is handled in the `pkg/config/config.yaml` file. 

All keys defined in the configuration file can be overriden using environment variables or command line arguments.

**e.g.**:

```console
go run main.go --aws.access_key_id=...
```

```env
AWS_ACCESS_KEY_ID=...
```