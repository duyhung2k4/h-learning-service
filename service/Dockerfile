FROM golang:1.22-alpine

WORKDIR /app

COPY cmd/core-service/ .

RUN go mod tidy
RUN go build -o main .

CMD ["./main"]
