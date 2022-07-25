PKG := github.com/xiaowenxia/nat
GOENV := GOPROXY=https://goproxy.io,direct GO111MODULE=on
GOOPTS := -mod=vendor

GOBUILD := $(GOENV) go build $(GOOPTS)
GOTEST := $(GOENV) go test $(GOOPTS)
AONE_APPNAME := nat

all: vendor build

vendor:
	@$(GOENV) go mod download
	@$(GOENV) go mod vendor

build:
	@$(GOBUILD) -o nat main.go

package: vendor build
	@tar czf ./nat.tar.gz ./nat ./statics/dist ./nat.json

test:
	@$(GOTEST) ./...

clean:
	@rm -rf nat
	@go clean -testcache