ARG GO_VERSION=1.23

FROM golang:${GO_VERSION} AS build
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app
COPY . .

RUN go get ./...
RUN go build -ldflags "-w -s" -o parser-service ./cmd/parser/main.go

FROM bitnami/minideb AS run
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app
COPY --from=build /app/parser-service /usr/bin/
RUN chmod +x /usr/bin/parser-service

ENTRYPOINT ["parser-service"]
