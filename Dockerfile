# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY ./docs ./docs


RUN go build -o main ./cmd/server
RUN go build -o seed ./scripts/seed.go


# Run stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/seed .

EXPOSE 8080

CMD ["./main"]
