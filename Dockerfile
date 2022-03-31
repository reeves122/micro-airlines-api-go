FROM golang:1.17 AS builder

WORKDIR /app

COPY ./src/go.mod .
COPY ./src/go.sum .

RUN go mod download

COPY ./src .

RUN go build -o micro-airlines-api-go ./cmd/


FROM debian:buster-slim

RUN apt-get update

COPY --from=builder /app/micro-airlines-api-go /

CMD ["/micro-airlines-api-go"]