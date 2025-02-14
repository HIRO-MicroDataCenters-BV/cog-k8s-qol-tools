FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/test-service

FROM alpine:latest

COPY --from=builder /app/test-service /test-service

EXPOSE 8080

CMD ["/test-service"]
