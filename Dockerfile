FROM golang:1.23.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env .env

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]