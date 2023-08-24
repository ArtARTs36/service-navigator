# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

ARG APP_VERSION="undefined"
ARG BUILD_TIME="undefined"

WORKDIR /go/src/github.com/artarts36/service-navigator

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -ldflags="-s -w" -o /go/bin/service-navigator /go/src/github.com/artarts36/service-navigator/cmd/main.go

# https://github.com/opencontainers/image-spec/blob/main/annotations.md
LABEL org.opencontainers.image.title="service-navigator"
LABEL org.opencontainers.image.description="navigator for your local docker projects in single network"
LABEL org.opencontainers.image.url="https://github.com/artarts36/service-navigator"
LABEL org.opencontainers.image.source="https://github.com/artarts36/service-navigator"
LABEL org.opencontainers.image.vendor="ArtARTs36"
LABEL org.opencontainers.image.version="$APP_VERSION"
LABEL org.opencontainers.image.created="$BUILD_TIME"
LABEL org.opencontainers.image.licenses="MIT"

EXPOSE 8080

CMD ["/go/bin/service-navigator"]
