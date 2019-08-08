# Go related variables.
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)/cmd:$(GOBASE)/pkg
GOBIN := $(GOBASE)/bin
GOFILES := $(shell go list ./...)

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## build: get missing dependencies and build all
build: go-mod-tidy go-install

## test: run test cases
test: go-test

## exec: run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
exec:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(run)

## clean: clean build files. Runs `go clean` internally.
clean:
	@-rm $(GOBIN) 2> /dev/null
	@-$(MAKE) go-clean

.PHONY: help
help: Makefile
	@echo
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

go-install:
	@echo "  >  Building binaries..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

go-mod-tidy:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod tidy

go-test:
	@echo "  >  Run test cases..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean $(GOFILES)
