all: bin/example

PLATFORM=local

.PHONY: bin/example
bin/example: go.sum
	@docker build . --target bin --output bin/ --platform ${PLATFORM}

go.mod:
	@docker run --rm -v ${CURDIR}:/src -w /src golang:1.14.3-alpine go mod init example

go.sum: go.mod
	@docker run --rm -v ${CURDIR}:/src -w /src golang:1.14.3-alpine go mod tidy