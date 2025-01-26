FROM alpine:latest
ARG TARGETARCH

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY ./bin/server-${TARGETARCH} /server

CMD ["/server"]