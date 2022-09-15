FROM golang:1.14.3-alpine AS base

# Install GCC
RUN apk add build-base
RUN mkdir /deps

WORKDIR /src

ENV CGO_ENABLED=0
ENV GOPATH=/

COPY ./lib /deps
COPY ./sources .
COPY ./tests .

FROM base AS build

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=gcc go build -o /out/server .

FROM base AS unit-test
ENTRYPOINT ["go", "test"]
CMD ["-v", "./..."]

FROM alpine:latest

RUN mkdir /config
RUN mkdir /install
RUN mkdir /web

COPY --from=build /out/server /install/server
COPY ./web /web
COPY VERSION /config/VERSION

CMD ["/install/server"]
