FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

RUN go mod download

# Copy entire repository
COPY . .

# Build Product Service
RUN CGO_ENABLED=0 GOOS=linux go build \
    -o product-service \
    ./services/product-service/cmd/server

# ---------------- Runtime ----------------

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/product-service .

EXPOSE 8081

CMD ["./product-service"]