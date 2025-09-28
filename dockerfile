# Development stage
FROM golang:alpine AS dev
WORKDIR /app
COPY --from=cosmtrek/air:v1.61.7 /go/bin/air  /go/bin/air
COPY . .
ENTRYPOINT ["/go/bin/air"]

# Production stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]