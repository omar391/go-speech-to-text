FROM alpine:latest
WORKDIR /build
COPY . .
RUN rm -fr ./.history ./.github ./git ./.vscode ./web/.parcel-cache
RUN apk add --no-cache  ca-certificates go ffmpeg
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GO111MODULE=on

ENV GIN_MODE=release
RUN go build -v -o /stt-service .
ENTRYPOINT ["./stt-service"]