APP_NAME=go-server

all: bin

.PHONY: dev
dev: bin
	@docker run --rm -d -p 4242:80 --name=$(APP_NAME)-dev -e PORT=80 $(APP_NAME)

.PHONY: clean
clean:
	@docker stop $(APP_NAME)-dev

.PHONY: bin
bin: go.sum
	@docker build . -t $(APP_NAME)

.PHONY: unit-test
build-test: go.sum
	@docker build . --target unit-test -t $(APP_NAME)-test
unit-test: build-test
	@docker run --rm $(APP_NAME)-test

go.mod:
	@docker run --rm -v ${CURDIR}:/src -w /src golang:1.14.3-alpine go mod init example

go.sum: go.mod
	@docker run --rm -v ${CURDIR}:/src -w /src golang:1.14.3-alpine go mod tidy