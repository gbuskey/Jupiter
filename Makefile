GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=jupiter
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_DESTINATION=./bin

all: build
build:
	$(GOBUILD) -o $(BINARY_DESTINATION)/$(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -r $(BINARY_DESTINATION)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin/$(BINARY_UNIX) -v

docker-build:

