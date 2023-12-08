FROM golang:1.21-alpine3.18@sha256:5c1cabd9a3c6851a3e18735a2c133fbd8f67fe37eb3203318b7af2ffd2547095 AS builder

WORKDIR /bubbles
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -trimpath ./


FROM alpine:latest@sha256:b295f5f1fa78532ebd27af26ea7338d554ae09b1d8db4a22b6b060a264f0e9cf AS runner

WORKDIR /home/bubbles
COPY --from=builder /bubbles/bot ./
EXPOSE 8080
CMD ["./bot"]
