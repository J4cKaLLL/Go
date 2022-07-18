BUILDPATH=$(CURDIR)
GOPATH=${shell go env GOPATH}
GO=$(shell which go)
GOINSTALL=$(GO) install
BINARY=customer
TESTS=go test $$(go list ./... | grep -v /vendor/) -cover

makedir:
	@if [ ! -d $(BUILDPATH)/build/bin ] ; then mkdir -p $(BUILDPATH)/build/bin ; fi
	@if [ ! -d $(BUILDPATH)/internal/entrypoint/api/grpc/proto ] ; then mkdir -p $(BUILDPATH)/internal/entrypoint/api/grpc/proto ; fi

proto: makedir
	@echo "Generate proto files..."
	@protoc -I api/v1/proto \
		api/v1/proto/customer.proto \
		--go_out=plugins=grpc:internal/entrypoint/api/grpc/proto

	@protoc -I internal \
		internal/core/repository/find_customer/proto/find_customer.proto \
		--go_out=plugins=grpc:internal

build: makedir proto
	${TESTS}
	@go build -ldflags "-s -w" -o $(BUILDPATH)/build/bin/${BINARY} cmd/customer/main.go
	@echo "Binario generado en build/bin/customer"

build-images:
	@echo "Building images..."
	@eval $(minikube docker-env)
	@docker build -t obtener-cliente-dom -f iaas/docker/Dockerfile .
	@sh iaas/docker/tag-images.sh

deploy: build-images
	@echo "Deploy on Kubernetes..."
	@sh iaas/kubernetes/set-tag-images.sh
	@kubectl -n demo apply -f iaas/kubernetes/config-map.yaml
	@kubectl -n demo apply -f iaas/kubernetes/deploy.yaml

test: proto
	@go test ./... --coverprofile coverfile >> /dev/null
	@go tool cover -func coverfile
	@go tool cover -html=coverfile -o coverfile.html

sonar: clean proto
	@go test -short -coverprofile=coverfile `go list ./...|grep -v vendor/ |grep -v mocks` >> /dev/null
	@sonar-scanner \
  -Dsonar.projectKey=Golang-Domain-Test \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9000 \
  -Dsonar.login=57d26e8738da99a623cc7c474266ef9e5d702270

clean:
	@echo "Cleaning Project..."
	@if [ -d ${BUILDPATH}/build/bin ] ; then rm -rf $(BUILDPATH)/build/bin/${BINARY} ; fi
	@find . -name "*pb.go" -exec rm -rf {} \; 


.PHONY: clean install unittest proto build build-images deploy test sonar
