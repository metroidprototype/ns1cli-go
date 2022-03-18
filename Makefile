
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

build.%:
	[ -d dist ] || mkdir dist
	for ARCH in amd64 arm64; do \
		GOOS=$* GOARCH=$${ARCH} $(BUILD_CMD) $(LD_OPTS) -o dist/$(NAME)-$*-$${ARCH} .; \
	done


build-all: build.darwin build.linux build.windows
	@mv dist/ns1-windows-amd64 dist/ns1-windows-amd64.exe
	@mv dist/ns1-windows-arm64 dist/ns1-windows-arm64.exe

test: deps
	go test -coverprofile=coverage.out -cover $(shell go list ./... | xargs)
	go tool cover -html=coverage.out
