# Tahap Build
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o main .

# Tahap Deploy
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
RUN mkdir -p /app/uploads
CMD ["./main"]