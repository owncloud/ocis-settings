SHELL := bash
NAME := ocis-settings
IMPORT := github.com/owncloud/$(NAME)
BIN := bin
DIST := dist
HUGO := hugo
PROTO_VERSION := v0
PROTO_SRC := pkg/proto/$(PROTO_VERSION)

ifeq ($(OS), Windows_NT)
	EXECUTABLE := $(NAME).exe
	UNAME := Windows
else
	EXECUTABLE := $(NAME)
	UNAME := $(shell uname -s)
endif

ifeq ($(UNAME), Darwin)
	GOBUILD ?= go build -i
else
	GOBUILD ?= go build
endif

PACKAGES ?= $(shell go list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f -not -path "./node_modules/*")
GENERATE ?= $(PACKAGES)

TAGS ?=

ifndef OUTPUT
	ifneq ($(DRONE_TAG),)
		OUTPUT ?= $(subst v,,$(DRONE_TAG))
	else
		OUTPUT ?= testing
	endif
endif

ifndef VERSION
	ifneq ($(DRONE_TAG),)
		VERSION ?= $(subst v,,$(DRONE_TAG))
	else
		VERSION ?= $(shell git rev-parse --short HEAD)
	endif
endif

ifndef DATE
	DATE := $(shell date -u '+%Y%m%d')
endif

LDFLAGS += -s -w -X "$(IMPORT)/pkg/version.String=$(VERSION)" -X "$(IMPORT)/pkg/version.Date=$(DATE)"
DEBUG_LDFLAGS += -X "$(IMPORT)/pkg/version.String=$(VERSION)" -X "$(IMPORT)/pkg/version.Date=$(DATE)"
GCFLAGS += all=-N -l

.PHONY: all
all: build

.PHONY: sync
sync:
	go mod download

.PHONY: clean
clean:
	go clean -i ./...
	rm -rf $(BIN) $(DIST) $(HUGO)

.PHONY: fmt
fmt:
	gofmt -s -w $(SOURCES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

# .PHONY: staticcheck
# staticcheck:
# 	go run honnef.co/go/tools/cmd/staticcheck -tags '$(TAGS)' $(PACKAGES)

.PHONY: lint
lint:
	for PKG in $(PACKAGES); do go run golang.org/x/lint/golint -set_exit_status $$PKG || exit 1; done;

.PHONY: generate
generate:
	go generate $(GENERATE)

.PHONY: changelog
changelog:
	go run github.com/restic/calens >| CHANGELOG.md

.PHONY: test
test:
	go run github.com/haya14busa/goverage -v -coverprofile coverage.out $(PACKAGES)

.PHONY: install
install: $(SOURCES)
	go install -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' ./cmd/$(NAME)

.PHONY: build
build: $(BIN)/$(EXECUTABLE) $(BIN)/$(EXECUTABLE)-debug

$(BIN)/$(EXECUTABLE): $(SOURCES)
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./cmd/$(NAME)

$(BIN)/$(EXECUTABLE)-debug: $(SOURCES)
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(DEBUG_LDFLAGS)' -gcflags '$(GCFLAGS)' -o $@ ./cmd/$(NAME)

.PHONY: release
release: release-dirs release-linux release-windows release-darwin release-copy release-check

.PHONY: release-dirs
release-dirs:
	mkdir -p $(DIST)/binaries $(DIST)/release

.PHONY: release-linux
release-linux: release-dirs
	go run github.com/mitchellh/gox -tags 'netgo $(TAGS)' -ldflags '-extldflags "-static" $(LDFLAGS)' -os 'linux' -arch 'amd64 386 arm64 arm' -output '$(DIST)/binaries/$(EXECUTABLE)-$(OUTPUT)-{{.OS}}-{{.Arch}}' ./cmd/$(NAME)

.PHONY: release-windows
release-windows: release-dirs
	go run github.com/mitchellh/gox -tags 'netgo $(TAGS)' -ldflags '-extldflags "-static" $(LDFLAGS)' -os 'windows' -arch 'amd64' -output '$(DIST)/binaries/$(EXECUTABLE)-$(OUTPUT)-{{.OS}}-{{.Arch}}' ./cmd/$(NAME)

.PHONY: release-darwin
release-darwin: release-dirs
	go run github.com/mitchellh/gox -tags 'netgo $(TAGS)' -ldflags '$(LDFLAGS)' -os 'darwin' -arch 'amd64' -output '$(DIST)/binaries/$(EXECUTABLE)-$(OUTPUT)-{{.OS}}-{{.Arch}}' ./cmd/$(NAME)

.PHONY: release-copy
release-copy:
	$(foreach file,$(wildcard $(DIST)/binaries/$(EXECUTABLE)-*),cp $(file) $(DIST)/release/$(notdir $(file));)

.PHONY: release-check
release-check:
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/release/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

.PHONY: release-finish
release-finish: release-copy release-check

.PHONY: docs-copy
docs-copy:
	mkdir -p $(HUGO); \
	mkdir -p $(HUGO)/content/extensions; \
	cd $(HUGO); \
	git init; \
	git remote rm origin; \
	git remote add origin https://github.com/owncloud/owncloud.github.io; \
	git fetch; \
	git checkout origin/source -f; \
	rsync --delete -ax ../docs/ content/extensions/$(NAME)

.PHONY: docs-build
docs-build:
	cd $(HUGO); hugo

.PHONY: docs
docs: docs-copy docs-build

.PHONY: watch
watch:
	go run github.com/cespare/reflex -c reflex.conf

$(GOPATH)/bin/protoc-gen-go:
	GO111MODULE=off go get -v github.com/golang/protobuf/protoc-gen-go

$(GOPATH)/bin/protoc-gen-micro:
	GO111MODULE=on go get -v github.com/micro/protoc-gen-micro/v2

$(GOPATH)/bin/protoc-gen-microweb:
	GO111MODULE=off go get -v github.com/owncloud/protoc-gen-microweb

$(GOPATH)/bin/protoc-gen-swagger:
	GO111MODULE=off go get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

$(PROTO_SRC)/settings.pb.go: $(PROTO_SRC)/settings.proto
	protoc \
		-I=third_party/ \
		-I=$(PROTO_SRC)/ \
		--go_out=. settings.proto

$(PROTO_SRC)/settings.pb.micro.go: $(PROTO_SRC)/settings.proto
	protoc \
		-I=third_party/ \
		-I=$(PROTO_SRC)/ \
		--micro_out=. settings.proto

$(PROTO_SRC)/settings.pb.web.go: $(PROTO_SRC)/settings.proto
	protoc \
		-I=third_party/ \
		-I=$(PROTO_SRC)/ \
		--microweb_out=. settings.proto

$(PROTO_SRC)/settings.swagger.json: $(PROTO_SRC)/settings.proto
	protoc \
		-I=third_party/ \
		-I=$(PROTO_SRC)/ \
		--swagger_out=$(PROTO_SRC) settings.proto

.PHONY: protobuf
protobuf: $(GOPATH)/bin/protoc-gen-go $(GOPATH)/bin/protoc-gen-micro $(GOPATH)/bin/protoc-gen-microweb $(GOPATH)/bin/protoc-gen-swagger \
		  $(PROTO_SRC)/settings.pb.go $(PROTO_SRC)/settings.pb.micro.go $(PROTO_SRC)/settings.pb.web.go $(PROTO_SRC)/settings.swagger.json
