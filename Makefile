GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.16","$(shell printf "$(GO_VERSION_SHORT)\n1.16" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.16. Found: $(GO_VERSION_SHORT))
endif

export GO111MODULE=on

SERVICE_NAME=ocp-template-api
SERVICE_PATH=ozoncp/ocp-template-api

PGV_VERSION:="v0.6.1"
GOOGLEAPIS_VERSION="master"
BUF_VERSION:="v0.51.0"
GOBIN?=$(GOPATH)/bin

.PHONY: run
grpc-server:
	go run cmd/grpc-server/main.go

.PHONY: lint
lint:
	golangci-lint run ./...


.PHONY: test
test:
	go test -v -race -timeout 30s -coverprofile cover.out ./...
	go tool cover -func cover.out | grep total | awk '{print $3}'


# ----------------------------------------------------------------

.PHONY: generate
generate: .generate

.generate:
	@command -v buf 2>&1 > /dev/null || (mkdir -p $(GOBIN) && curl -sSL0 https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-$(shell uname -s)-$(shell uname -m) -o $(GOBIN)/buf && chmod +x $(GOBIN)/buf)
	PATH=$(GOBIN):$(PATH) buf generate
	cd pkg/$(SERVICE_NAME) && ls go.mod || go mod init github.com/$(SERVICE_PATH)/pkg/$(SERVICE_NAME) && go mod tidy

# ----------------------------------------------------------------

.PHONY: deps
deps: .deps .bin-deps

.deps:
	@[ -f go.mod ] || go mod init github.com/$(SERVICE_PATH)
	find . -name go.mod -exec bash -c 'pushd "$${1%go.mod}" && go mod tidy && popd' _ {} \;

.bin-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/envoyproxy/protoc-gen-validate@$(PGV_VERSION)
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: build
build: generate
		go mod download && CGO_ENABLED=0  go build \
			-tags='no_mysql no_sqlite3' \
			-ldflags=" \
				-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
				-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
			" \
			-o ./bin/grpc-server ./cmd/grpc-server/main.go
