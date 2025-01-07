# syntax=docker/dockerfile:1
FROM golang:1.23 AS builder
WORKDIR /src
COPY . .
RUN go build -a -installsuffix cgo -o /bin/proxy ./cmd/

FROM scratch
COPY --from=builder /bin/proxy /bin/proxy
CMD ["/bin/proxy"]
