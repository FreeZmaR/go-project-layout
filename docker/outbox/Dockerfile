FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build cmd/outbox -o app

CMD ["./app"]