ARG GO_VERSION=1.23

FROM golang:${GO_VERSION} AS build
WORKDIR /app
COPY . .

RUN go get ./...
RUN go build -ldflags "-w -s" -o scraper-service cmd/scraper/main.go

FROM alpine:latest AS run
WORKDIR /app
COPY --from=build /app/scraper-service /usr/bin/scraper-service
ENTRYPOINT ["scraper-service"]
