APP_NAME=toy

GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin

BIN_BUILD_FLAGS?=CGO_ENABLED=0

GOVVV=go run github.com/ahmetb/govvv@v0.3.0 
BIN_VERSION?="git"
GOVVV_FLAGS=$(shell $(GOVVV) -flags -version $(BIN_VERSION) -pkg $(shell go list ./buildinfo))

PROTOC_GEN_GO = $(GOBIN)/protoc-gen-go # this was installed previously using go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
BUF=go run github.com/bufbuild/buf/cmd/buf@v0.41.0

build: 
	$(BIN_BUILD_FLAGS) go build -ldflags="${GOVVV_FLAGS}" ./cmd/${APP_NAME}
.PHONY: build

protos: $(PROTOC_GEN_GO) clean-protos
	$(BUF) generate --template '{"version":"v1beta1","plugins":[{"name":"go","out":"gen","opt":"paths=source_relative","path":$(PROTOC_GEN_GO)}]}'
.PHONY: protos

clean-protos:
	find . -type f -name '*.pb.go' -delete
	find . -type f -name '*pb_test.go' -delete
.PHONY: clean-protos