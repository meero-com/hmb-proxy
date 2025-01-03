# syntax=docker/dockerfile:1
FROM golang:1.23 AS builder
WORKDIR /src
COPY . .
RUN go build -o /bin/dummy-service .

FROM debian:12
RUN apt update && apt install -y ca-certificates
COPY --from=builder /bin/dummy-service /bin/dummy-service
CMD ["/bin/dummy-service"]
