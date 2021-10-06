.DEFAULT_GOAL := build

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	$(shell go env GOPATH)/bin/golint ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: build
build: fmt
	go build ./...
	go build

.PHONY: clean
clean:
	rm animal-mongo

.PHONY: run
run: build
	./animal-mongo
