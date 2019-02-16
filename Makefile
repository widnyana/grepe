PROJECTNAME := $(shell basename "$(PWD)")
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)


MAKEFLAGS += --silent

install: go-get

## build: Build all binary.
build:
	@echo "  >  Building main binary"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)
	@echo "  >  Building gql binary"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/gql ./cmd/gql

## clean: Clean build files. Runs `go clean` internally.
clean:
	@-rm $(GOBIN)/$(PROJECTNAME) 2> /dev/null
	@-rm $(GOBIN)/gql 2> /dev/null
	@-$(MAKE) go-clean

go-build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)


go-compile: go-get go-build


go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	dep ensure

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

go-install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
