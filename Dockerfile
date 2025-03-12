FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o books-api ./cmd/api

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/books-api .

EXPOSE 5000

CMD ["./books-api"]
