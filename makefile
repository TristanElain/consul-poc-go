
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
		del /F $(GOBIN)\$(BINARY_NAME).exe
		del /F $(GOBIN)\$(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)

# Cross compilation
build-linux: export CGO_ENABLED=0 export GOOS=linux export GOARCH=amd64
build-linux: 
	$(GOBUILD) -o $(GOBIN)\$(BINARY_UNIX) -v -i $(GOPROJECT)