FROM golang:1.21-alpine3.18@sha256:70afe55365a265f0762257550bc38440e0d6d6b97020d3f8c85328f00200dd8e AS builder

WORKDIR /bubbles
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -trimpath ./


FROM alpine:latest@sha256:34871e7290500828b39e22294660bee86d966bc0017544e848dd9a255cdf59e0 AS runner

WORKDIR /home/bubbles
COPY --from=builder /bubbles/bot ./
EXPOSE 8080
CMD ["./bot"]
