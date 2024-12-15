ARG GO_VERSION=1.23

FROM golang:${GO_VERSION} AS build
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app
COPY . .

RUN go env -w CGO_ENABLED=1  # Enable CGO for sqlite3
RUN go get ./...
RUN go build -ldflags "-w -s" -o scraper-service cmd/scraper/main.go

FROM bitnami/minideb AS run
RUN apt-get update && apt-get install -y ca-certificates sqlite3
WORKDIR /app
COPY --from=build /app/scraper-service /usr/bin/
RUN chmod +x /usr/bin/scraper-service

ENTRYPOINT ["scraper-service"]
