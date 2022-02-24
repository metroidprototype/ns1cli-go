
NAME = ns1
COMMIT = $(shell git rev-parse --short HEAD)
BUILDTIME = $(shell date +%Y-%m-%dT%T%z)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
VERSION ?= $(BRANCH)-$(COMMIT)

BUILD_CMD = CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo
LD_OPTS = -ldflags="-X github.com/metroidprototype/ns1cli-go/version.Version=$(VERSION) -w"

clean:
	go clean -i ./...
	rm -rf ./dist

deps:
	go mod tidy

fmt:
	go fmt ./...
go test -coverprofile=coverage.out ./...
build.%:
	[ -d dist ] || mkdir dist
	GOOS=$* $(BUILD_CMD) $(LD_OPTS) -o dist/$(NAME)-$* .

build-all: build.darwin build.linux build.windows
	@mv dist/ns1-windows dist/ns1-windows.exe

test: deps
	go test -coverprofile=coverage.out -cover $(shell go list ./... | xargs)
	go tool cover -html=coverage.out
