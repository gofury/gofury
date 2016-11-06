default: deps test

lint:
	@go fmt

deps:
	@glide --quiet install

test:
	@go test -v -race $(shell glide nv)