# Tahap Build
FROM golang:1.24.1-alpine AS builder
WORKDIR /app

# Copy file go.mod dan go.sum, lalu download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build aplikasi
RUN go build -o main .

# Tahap Deploy (Image lebih ringan)
FROM alpine:latest
WORKDIR /root/

# Copy binary dari tahap build
COPY --from=builder /app/main .

# Set permission agar binary bisa dieksekusi
RUN chmod +x main

# Jalankan aplikasi
CMD ["./main"]
