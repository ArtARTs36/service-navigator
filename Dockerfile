# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /go/src/github.com/artarts36/potaynik

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -ldflags="-s -w" -o /go/bin/service-navigator /go/src/github.com/artarts36/potaynik/cmd/main.go

EXPOSE 8080

CMD ["/go/bin/service-navigator"]
