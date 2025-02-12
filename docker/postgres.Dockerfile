FROM postgres:15-alpine

ARG HOST_UID=1000
ARG HOST_GID=1000
ENV HOST_UID=${HOST_UID:-1000} \
    HOST_GID=${HOST_GID:-1000}

RUN set -eux \
    && apk add --update --no-cache shadow \
    && rm -rf /var/cache/apk/* \
    && rm -rf /tmp/*

RUN set -eux \
    && usermod -u ${HOST_UID} postgres \
    && groupmod -g ${HOST_GID} postgres

USER postgres

EXPOSE 5432
