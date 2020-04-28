GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=main
BINARY_UNIX=$(BINARY_NAME)_unix

build:
    $(GOBUILD) -o $(BINARY_NAME) -v
test:
    $(GOTEST) -v ./...
clean:
    $(GOCLEAN)
    rm -f $(BINARY_NAME)
    rm -f $(BINARY_UNIX)
run:
    $(GOBUILD) -o $(BINARY_NAME) -v ./...
    ./$(BINARY_NAME)
gomod:
    $(GOMOD) tidy


# クロスコンパイル
build-linux:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
lint:
    golint -set_exit_status $$(go list ./...)
    go vet ./...:
