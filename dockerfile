# Development stage
FROM golang:alpine AS dev
WORKDIR /app
COPY --from=cosmtrek/air:v1.61.7 /go/bin/air  /go/bin/air
COPY fizzbuzz/ .
ENTRYPOINT ["/go/bin/air"]

# Production stage
FROM golang:alpine AS builder
WORKDIR /app
COPY fizzbuzz/ .
RUN go mod download
RUN go build -o main ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]