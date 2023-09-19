FROM golang:1.21-alpine3.18@sha256:96634e55b363cb93d39f78fb18aa64abc7f96d372c176660d7b8b6118939d97b AS builder

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


#FROM busybox:1.36 AS debugger

#RUN mkdir /tmp/debug
#COPY --from=builder ./ /tmp/debug
# sudo docker run --rm -it git.toast-server.net/toast/bubbles:main find /tmp/debug/bubbles/ >> output.log