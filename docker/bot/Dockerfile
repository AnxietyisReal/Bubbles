FROM golang:1.21-alpine3.18@sha256:d7d6dd2825586ab8fdb83b66c3a056e79845b1712fa4739062451375b9de328c AS builder

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
