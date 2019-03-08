GORESTSECURITY_PKG_VERSION?=0.0.0
COMMIT=`git rev-parse --short HEAD`

build: 
	go build -v --ldflags "-w \
	-X github.com/mrcaelumn/Go-REST-API-Security/version.Version/version.Version=$(GORESTSECURITY_PKG_VERSION) \
	-X github.com/mrcaelumn/Go-REST-API-Security/version.Version/version.GitCommit=$(COMMIT)" .

build_binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Go-REST-API-Security -a --ldflags "-w \
	-X github.com/mrcaelumn/Go-REST-API-Security/version.Version=$(GORESTSECURITY_PKG_VERSION) \
	-X github.com/mrcaelumn/Go-REST-API-Security/version.GitCommit=$(COMMIT)" .

test:
	@go test -v $(shell go list ./... | grep -v /vendor/)

vet:
	@go vet -v $(shell go list ./... | grep -v /vendor/)

clean:
	@rm -rf build
	@rm -rf Go-REST-API-Security*

.PHONY: test vet build build_binary clean
