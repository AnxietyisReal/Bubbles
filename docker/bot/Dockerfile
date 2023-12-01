FROM golang:1.21-alpine3.18@sha256:ff7d5e6e679675001a75f0fafa0f3383d50fc802e9f40fcd5e6bf3fcab5e1d15 AS builder

WORKDIR /bubbles
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -trimpath ./


FROM alpine:latest@sha256:69e0a8b6d2048be90290c992957a3cde2e87e955e5bb9c27d228c78c46c7840f AS runner

WORKDIR /home/bubbles
COPY --from=builder /bubbles/bot ./
EXPOSE 8080
CMD ["./bot"]
