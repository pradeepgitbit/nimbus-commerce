# ---------- Builder Stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o product-service \
    ./services/product-service/cmd/server

# ---------- Runtime Stage ----------
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/product-service .

EXPOSE 8081

CMD ["./product-service"]