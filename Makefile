GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

GOQLINES := qlines

.PHONY: qlines
qlines:
	cd $(GOQLINES) && go build -o $(GOBIN)/$(GOQLINES) main.go	

build: qlines
all: build