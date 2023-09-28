FROM golang:1.21-alpine3.18@sha256:08d551a16476f47ac80fd9e5f3f128de66c46ededdf2a582a555353a53c578e5 AS builder

WORKDIR /bubbles
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -trimpath ./


FROM alpine:latest@sha256:7144f7bab3d4c2648d7e59409f15ec52a18006a128c733fcff20d3a4a54ba44a AS runner

WORKDIR /home/bubbles
COPY --from=builder /bubbles/bot ./
EXPOSE 8080
CMD ["./bot"]
