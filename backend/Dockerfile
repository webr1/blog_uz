FROM golang:1.24-alpine AS builder
WORKDIR /app
ENV GOTOOLCHAIN=auto

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6
COPY . .
RUN swag init -g cmd/main.go

RUN CGO_ENABLED=0 go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/env ./env

EXPOSE 8080
CMD ["./main"]