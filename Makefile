# ANSI color
RED=\033[31m
GREEN=\033[32m
RESET=\033[0m

COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

test:
	go test -v ./soumu/... | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

lint:
	golint --set_exit_status $$(go list ./...)
	go vet ./...

clean:
	rm -f example/get-number/get-number
	rm -f example/get-list/get-list
