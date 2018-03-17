# Go parameters
# for use in docker image with all GO env setup
# /go -> {src,bin,pkg}
GOCMD=go
GOBIN=/go/bin
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=myInit
BINARY_UNIX=$(BINARY_NAME)_unix
SRC_FILES=sandbox.go utils.go

all: test build
build:
	$(GOBUILD) -o ${GOBIN}/$(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f ${GOBIN}/$(BINARY_NAME)
	rm -f ${GOBIN}/$(BINARY_UNIX)
run:
	$(GORUN) ${SRC_FILES}
exec:
	$(GOBUILD) -o ${GOBIN}/$(BINARY_NAME) -v ./...
	/${GOBIN}/$(BINARY_NAME)
#deps:
#        $(GOGET) github.com/id/src1
#        $(GOGET) github.com/id/src2


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ${GOBIN}/$(BINARY_UNIX) -v
#docker-build:
#        docker run --rm -it -v "$(GOPATH)":/go -w /go/src/ golang:latest go build -o "$(BINARY_UNIX)" -v
