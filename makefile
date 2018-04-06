
# File variables
GOWORKSPACE=C:\workspace_go
GOPROJECT=consul-poc
PROJECT_PATH:=$(GOWORKSPACE)\$(GOPROJECT)
BINARY_NAME=consul-poc
BINARY_UNIX=$(BINARY_NAME).bin

# Environment variables
GOBIN=$(PROJECT_PATH)\bin
GOPATH:=$(PROJECT_PATH);$(GOPATH)
# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Tasks
all: test build
build: 
		$(GOINSTALL) -v $(GOPROJECT)
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(GOBIN)\$(BINARY_NAME).exe
		rm -f $(GOBIN)\$(BINARY_UNIX)
run:
		build
		start bin/$(BINARY_NAME)

# Cross compilation
build-linux: export CGO_ENABLED=0
build-linux: export GOARCH=arm
build-linux: export GOOS=linux
build-linux: 
	$(GOBUILD) -o $(GOBIN)\$(BINARY_UNIX) -v -i $(GOPROJECT)
	zip bin/templates.zip templates