GORESTSECURITY_PKG_VERSION?=0.0.0
COMMIT=`git rev-parse --short HEAD`

build-mac: 
	@GOOS=darwin GOARCH=amd64 \
	go build -v --ldflags "-w \
	-X github.com/mrcaelumn/go-rest-api-security/version.Version/version.Version=$(GORESTSECURITY_PKG_VERSION) \
	-X github.com/mrcaelumn/go-rest-api-security/version.Version/version.GitCommit=$(COMMIT)" .

build_binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-rest-api-security -a --ldflags "-w \
	-X github.com/mrcaelumn/go-rest-api-security/version.Version=$(GORESTSECURITY_PKG_VERSION) \
	-X github.com/mrcaelumn/go-rest-api-security/version.GitCommit=$(COMMIT)" .

test:
	@go test -v $(shell go list ./... | grep -v /vendor/)

vet:
	@go vet -v $(shell go list ./... | grep -v /vendor/)

clean:
	@rm -rf build
	@rm -rf go-rest-api-security*

.PHONY: test vet build build_binary clean
