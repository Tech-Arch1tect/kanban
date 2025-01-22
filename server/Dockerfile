FROM golang:1.23-alpine AS build

ENV CGO_ENABLED=1

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN go build -o /server cmd/main.go 

## Deploy
FROM alpine:latest

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY --from=build /server /server

CMD ["/server"]