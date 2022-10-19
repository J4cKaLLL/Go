BUILDPATH=$(CURDIR)
GOPATH=$(shell go env GOPATH)
GO=$(shell which go)
GOINSTALL=$(GO) install
BINARY=beer

makedir:
	@if [ ! -d $(BUILDPATH)/build/bin ] ; then mkdir -p $(BUILDPATH)/build/bin ; fi

build: makedir test
	@echo "Compilando..."
	@go build -ldflags "-s -w" -o $(BUILDPATH)/build/bin/${BINARY} cmd/${BINARY}/main.go
	@echo "Binario generado en build/bin/"${BINARY}

test:
	@echo "Ejecutando tests..."
	@go test ./... --coverprofile coverfile >> /dev/null
	@go tool cover -func coverfile
	@go tool cover -html=coverfile -o coverfile.html

clean:
	@echo "Limpiando binarios..."
	@if [ -d $(BUILDPATH)/build/bin ] ; then rm -rf $(BUILDPATH)/build/ ; fi
	@rm -rf coverfile*

benchmark:
	@echo "Ejecutando benchmarks..."
	@go test -benchmem ./... -bench=.

run: docker-build
	@echo "Running..."
	@docker run --rm -d -p 8080:8080 ${IMAGE}

docker-build:
	@echo "Building image..."
	@docker build -t ${IMAGE} -f iaas/Dockerfile .

.PHONY: build test run docker-build clean%