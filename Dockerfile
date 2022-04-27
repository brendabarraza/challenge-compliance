FROM golang:1.17 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o bin/challenge_service ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/credentials.json .
COPY --from=builder /app/dev.env .
WORKDIR /app/bin
COPY --from=builder /app/bin/challenge_service .
CMD ["./challenge_service"]