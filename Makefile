setup:
	go get \
		github.com/goreleaser/goreleaser \
		github.com/golang/lint/golint

test: setup
	go test -v ./...

lint: setup
	go vet ./...
	golint -set_exit_status ./...

.PHONY: setup test lint
