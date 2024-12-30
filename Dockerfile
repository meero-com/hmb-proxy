# syntax=docker/dockerfile:1
FROM golang:1.23 AS builder
WORKDIR /src
COPY . .
RUN go build -o /bin/proxy ./proxy/

FROM scratch
COPY --from=builder /bin/proxy /bin/proxy
CMD ["/bin/proxy"]
