FROM golang:1.23 AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags "-s -w" -o hmb-proxy ./cmd/main.go

# FROM scratch
# COPY --from=builder /src/main /src/main

ENTRYPOINT ["/src/hmb-proxy"]
