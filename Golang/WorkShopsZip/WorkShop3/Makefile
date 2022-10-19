BUILDPATH=$(CURDIR)
GOPATH=${shell go env GOPATH}
GO=$(shell which go)
GOINSTALL=$(GO) install
BINARY=example_beers
TESTS=go test $$(go list ./... | grep -v /vendor/) -cover

makedir:
	@if [ ! -d $(BUILDPATH)/build/bin ] ; then mkdir -p $(BUILDPATH)/build/bin ; fi

build: 
	${TESTS}
	@go build -ldflags "-s -w" -o $(BUILDPATH)/build/bin/${BINARY} cmd/example_auth/main.go
	@echo "Binario generado en build/bin/example_auth"

test: 
	@go test ./... --coverprofile coverfile >> /dev/null
	@go tool cover -func coverfile
	@go tool cover -html=coverfile -o coverfile.html

clean:
	@echo "Cleaning Project..."
	@if [ -d ${BUILDPATH}/build/bin ] ; then rm -rf $(BUILDPATH)/build/bin/${BINARY} ; fi

.PHONY: build test clean build-images deploy