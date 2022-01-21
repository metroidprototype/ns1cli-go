
NAME = ns1
COMMIT = $(shell git rev-parse --short HEAD)
BUILDTIME = $(shell date +%Y-%m-%dT%T%z)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
VERSION ?= $(BRANCH)-$(COMMIT)

BUILD_CMD = CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo
LD_OPTS = -ldflags="-X github.com/metroidprototype/ns1cli-go/VERSION=$(VERSION) -w"

clean:
	go clean -i ./...
	rm -rf ./dist

deps:
	go mod tidy

fmt:
	go fmt ./...

build.%:
	[ -d dist ] || mkdir dist
	GOOS=$* $(BUILD_CMD) $(LD_OPTS) -o dist/$(NAME)-$* .


test: deps
	go test -cover $(shell go list ./... | xargs)
