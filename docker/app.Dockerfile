FROM golang:1.24.3-alpine

ARG HOST_UID=1000
ARG HOST_GID=1000
ENV HOST_UID=${HOST_UID:-1000} \
    HOST_GID=${HOST_GID:-1000}

RUN go install github.com/githubnemo/CompileDaemon@latest

RUN set -eux &&\
    apk add --update --no-cache shadow bash &&\
    addgroup -g ${HOST_GID} -S appuser &&\
    adduser -u ${HOST_UID} -D -S -G appuser appuser && exit 0 ; exit 1 &&\
    rm -rf /var/cache/apk/* &&\
    rm -rf /tmp/*

WORKDIR /app
COPY ./ /app

RUN go mod download

EXPOSE 7100

ENTRYPOINT CompileDaemon \
    --exclude-dir=.git \
    --exclude-dir=docker \
    --build="go build -o /boilerplate-http ./cmd/http/main.go" \
    --command=/boilerplate-http
