
.PHONY: build

configure:
	@cmake -S . -B _build

build: configure
	@cmake --build _build

test: build
	@ctest --test-dir _build --output-on-failure --progress