FROM golang:1.14.3-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0

COPY ./sources .
COPY ./tests .

FROM base AS build
RUN GOOS=linux GOARCH=amd64 go build -o /out/server .

FROM base AS unit-test
ENTRYPOINT ["go", "test"]
CMD ["-v", "."]

FROM alpine:latest

RUN mkdir /install
RUN mkdir /web

COPY --from=build /out/server /install/server
COPY ./web /web

CMD ["/install/server"]
