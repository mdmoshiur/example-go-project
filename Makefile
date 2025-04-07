export APP=example-go
export EXAMPLE_GO_ENV=development
export EXAMPLE_GO_CONSUL_URL=127.0.0.1:8500
export EXAMPLE_GO_CONSUL_PATH=$(APP)
.PHONY: all test coverage
build-run:
	go run . serve -v

migration-up:
	go build -v .
	./example-go migration up -p
migration-down:
	go build -v .
	./example-go migration down
migration-reset:
	go build -v .
	./example-go migration reset
migration-seed-all:
	go build -v .
	./example-go migration seed all

format:
	gofmt -l -s -w .
get:
	go get ./...
build:
	go build ./...
install:
	go install ./...
lint:
	golangci-lint run ./...

test:
	go test ./... -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt

coverage: test
	go tool cover -html=.coverage.txt

