FROM golang:1.21-alpine3.18@sha256:110b07af87238fbdc5f1df52b00927cf58ce3de358eeeb1854f10a8b5e5e1411 AS builder

WORKDIR /bubbles
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -trimpath ./


FROM alpine:latest@sha256:3a244917cd8d31893d88232d2201755fc9c884fde57a47d148f77edbcc8ad2a6 AS runner

WORKDIR /home/bubbles
COPY --from=builder /bubbles/bot ./
EXPOSE 8080
CMD ["./bot"]
