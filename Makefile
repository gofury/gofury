Mdefault: deps test

lint:
	@go fmt
.PHONY: lint

deps:
	@glide --quiet install
.PHONY: deps

test:
	@go test -v -race $(shell glide nv)
.PHONY: test