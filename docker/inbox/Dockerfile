FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build cmd/inbox -o app

CMD ["./app"]