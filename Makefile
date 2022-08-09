all: bin

.PHONY: dev
dev: bin
	@docker run --rm -d -p 4242:80 --name=server-dev -e PORT=80 go-server

.PHONY: bin
bin: go.sum
	@docker build . -t go-server

.PHONY: unit-test
build-test: go.sum
	@docker build . --target unit-test -t go-server-test
unit-test: build-test
	@docker run --rm example-test

go.mod:
	@docker run --rm -v ${CURDIR}:/src -w /src golang:1.14.3-alpine go mod init example

go.sum: go.mod
	@docker run --rm -v ${CURDIR}:/src -w /src golang:1.14.3-alpine go mod tidy