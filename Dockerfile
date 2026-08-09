# Build stage
FROM golang:1.26.5-alpine3.24 AS builder
WORKDIR /app
COPY tellmi-sdk /tellmi-sdk
COPY tellmi-gateway/go.mod tellmi-gateway/go.sum ./
RUN go mod download
COPY tellmi-gateway/. .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

# Run stage
FROM alpine:3.21
RUN apk add --no-cache ca-certificates
RUN adduser -D -u 10001 appuser
WORKDIR /app
COPY --from=builder /app/main .
USER appuser
ENV ENVIRONMENT=docker
EXPOSE 8080
CMD ["/app/main"]
