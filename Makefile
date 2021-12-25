RED=\033[31m
GREEN=\033[32m
RESET=\033[0m

COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

DOCKER=docker-compose run --rm app

test:
	$(DOCKER) go test -v ./... | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

fmt:
	$(DOCKER) gofmt -s -l `find . -type f -name '*.go'`

vet:
	$(DOCKER) go list ./... | xargs go vet

up:
	docker-compose up -d
