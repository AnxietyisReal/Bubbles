FROM golang:1.21-alpine3.18@sha256:1c9cc949513477766da12bfa80541c4f24957323b0ee00630a6ff4ccf334b75b AS builder

WORKDIR /bubbles
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -trimpath ./


FROM alpine:latest@sha256:eece025e432126ce23f223450a0326fbebde39cdf496a85d8c016293fc851978 AS runner

WORKDIR /home/bubbles
COPY --from=builder /bubbles/bot ./
EXPOSE 8080
CMD ["./bot"]
