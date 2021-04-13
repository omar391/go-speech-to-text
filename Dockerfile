FROM alpine:latest
WORKDIR /build
COPY . .
RUN rm -fr ./.history ./.github ./git ./.vscode ./web/.parcel-cache ./go ./web/node_modules
RUN apk add --no-cache  ca-certificates go
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GO111MODULE=on
RUN go build -v -o /stt-service .


#doing multi-stage build (for ffmpeg and static app)
FROM alpine:latest
WORKDIR /
ENV GIN_MODE=release
RUN apk add --no-cache ffmpeg
COPY --from=0 /stt-service .
COPY /conf.d ./conf.d
COPY /data ./data
COPY /web ./web
ENTRYPOINT ["./stt-service"]