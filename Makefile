.DEFAULT_GOAL := check

.PHONY: check
check: clean fmt lint test

.PHONY: clean
clean: 
	rm -rf pkg/models/wScan.db

.PHONY: fmt
fmt:
	gofumpt -w pkg/ cmd/
	goimports -w -local github.com/digital-technology-agency/web-scan pkg/ cmd/

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	@ ! gofumpt -l pkg/ cmd/ | read || (echo "Bad format, run make fmt" && exit 1)
	go test -race -timeout 5s ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vendor
vendor: tidy
	go mod vendor
