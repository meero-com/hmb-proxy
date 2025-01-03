# Dummy test service

![architecture.png](./dummy-service-architecture.png)

## Build

**Go executable:**

```console
$ go build
$ ./dummy-service
```

**Container image:**

```console
$ docker build -t dummy:local .
$ docker run -i -t dummy:local
```

## Test the service

Start top-level localstack container:
```console
$ pushd ..
$ docker compose up -d localstack
$ popd
```

Send an SQS Message
```console
$ bash sqs-send.sh
```
