GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

GOQLINES := qlines
GOASTEROIDS := asteroids

.PHONY: qlines asteroids
qlines_build:
	cd $(GOQLINES) && go build -o $(GOBIN)/$(GOQLINES) main.go && cd $(GOBASE)
asteroids_build:
	cd $(GOASTEROIDS) && go build -o $(GOBIN)/$(GOASTEROIDS) main.go && cd $(GOBASE)	

qlines: qlines_build
	$(GOBIN)/$(GOQLINES)

asteroids: asteroids_build
	$(GOBIN)/$(GOASTEROIDS)
	
build: qlines_build asteroids_build

all: build