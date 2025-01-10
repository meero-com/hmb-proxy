# syntax=docker/dockerfile:1
FROM golang:1.23 AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

# FROM scratch
# COPY --from=builder /src/main /src/main

CMD ["/src/main"]
