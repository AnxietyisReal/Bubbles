FROM golang:1.21.1-alpine3.17 AS builder

WORKDIR /bubbles
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -trimpath ./


FROM alpine:latest AS runner

WORKDIR /home/bubbles
COPY --from=builder /bubbles/bot ./
CMD ["./bot"]


#FROM busybox:1.36 AS debugger

#RUN mkdir /tmp/debug
#COPY --from=builder ./ /tmp/debug
# sudo docker run --rm -it git.toast-server.net/toast/bubbles:main find /tmp/debug/bubbles/ >> output.log