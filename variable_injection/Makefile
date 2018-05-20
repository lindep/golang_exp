
GOBIN=/go/bin

.PHONY: all clean

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

BINARY=varInject
VERSION=0.1.1
#BUILD=$(shell git rev-parse HEAD)
BUILD=7e36207d046293b3b17067f1b
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

all:
	go build -o ${GOBIN}/$(BINARY) $(LDFLAGS)

clean:
	-rm ${GOBIN}/$(BINARY)
