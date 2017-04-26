###################
# CI ENTRY POINTS #
###################
test:
	docker-compose down
	docker-compose run golang make _deps _test
	docker-compose down
.PHONY: test

tag:
	git tag "$VERSION-b$GO_PIPELINE_COUNTER"
	git push origin "$VERSION-b$GO_PIPELINE_COUNTER"
.PHONY: tag

##########
# OTHERS #
##########
_lint:
	go fmt

_deps:
	glide --quiet install

_test:
	go test -v $(shell glide nv)