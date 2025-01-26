FROM debian:stable-slim
ARG TARGETARCH

WORKDIR /app

COPY ./bin/server-${TARGETARCH} /server

CMD ["/server"]