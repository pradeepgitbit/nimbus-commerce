FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
-o api-gateway \
./services/api-gateway/cmd/server

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/api-gateway .

EXPOSE 8083

CMD ["./api-gateway"]