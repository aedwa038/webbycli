BINARY_NAME=webby
CONFIG_NAME=config.yaml

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

build: dep
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin 
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux 
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

dep:
	go mod download

fmt:
	@gofmt -l -w $(SRC)

test:
	go test ./...

cov:
	go test ./... -coverprofile=coverage.out

package: build
	tar -cvf ${BINARY_NAME}-darwin.tar ${BINARY_NAME}-darwin $(CONFIG_NAME)
	tar -cvf ${BINARY_NAME}-linux.tar ${BINARY_NAME}-linux $(CONFIG_NAME)
	tar -cvf ${BINARY_NAME}-windows.tar ${BINARY_NAME}-windows $(CONFIG_NAME)
