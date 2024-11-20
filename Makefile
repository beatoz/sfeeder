ifeq ($(OS),Windows_NT)
	HOSTOS=windows
	ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
		HOSTARCH=amd64
	else
		ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
			HOSTARCH=amd64
		else ifeq ($(PROCESSOR_ARCHITECTURE),x86)
			HOSTARCH=amd32
		endif
	endif
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		HOSTOS=linux
	else ifeq ($(UNAME_S),Darwin)
		HOSTOS=darwin
	else
		@echo "Unknown OS: $(UNAME_S)"
		exit
	endif

	UNAME_M := $(shell uname -m)
	ifeq ($(UNAME_M),x86_64)
		HOSTARCH=amd64
	else ifneq ($(filter %86,$(UNAME_M)),)
		HOSTARCH=amd32
	else ifeq ($(UNAME_M),arm64)
		HOSTARCH=arm64
	else ifneq ($(filter arm%,$(UNAME_M)),)
		HOSTARCH=arm
	endif
endif

TARGETOS=$(HOSTOS)
ifdef MAKECMDGOALS
	ifeq ($(MAKECMDGOALS), $(filter $(MAKECMDGOALS), windows linux darwin))
		TARGETOS=$(MAKECMDGOALS)
	endif
endif

#GITTAG=$(shell git describe --tags $(shell git rev-list --tags --max-count=1))
GITCOMMIT=$(shell git log -1 --pretty=format:"%h")
BUILD_FLAGS=-a -ldflags "-w -s"

LOCAL_GOPATH = $(shell go env GOPATH)
BUILDDIR="./build/$(HOSTOS)"

all: pbm $(TARGETOS)

$(TARGETOS):
	@echo Build sfeeder for $(@) on $(UNAME_S)-$(UNAME_M)
ifeq ($(HOSTOS), windows)
	@set GOOS=$@& set GOARCH=$(HOSTARCH)& go build -o $(BUILDDIR)/sfeeder.exe $(BUILD_FLAGS)  ./cmd/
else
	@GOOS=$@ GOARCH=$(HOSTARCH) go build -o $(BUILDDIR)/sfeeder $(BUILD_FLAGS) ./sfeeder.go
endif

install:
	@echo "Install binaries to $(LOCAL_GOPATH)/bin"
	@cp $(BUILDDIR)/* $(LOCAL_GOPATH)/bin
pbm:
	@echo Compile protocol messages
	@protoc --go_out=$(LOCAL_GOPATH)/src --go-grpc_out=$(LOCAL_GOPATH)/src -I./protos secret_feeder.proto
