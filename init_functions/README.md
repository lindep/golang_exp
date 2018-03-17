


``` bash
lindep @ ~/projects/golang
└─ $ ls
bin  pkg  src
lindep @ ~/projects/golang

└─ $ docker run --rm -v $(pwd):/go -it golang bash
$ export GOBIN=$GOPATH/bin
$ go run sandbox.go utils.go
1
2
3
$ go build sandbox.go utils.go
$ ls
sandbox  sandbox.go  utils.go
$ ./sandbox
1
2
3
$ go install sandbox.go utils.go
$ ls $GOPATH/bin
sandbox
$ /$GOPATH/bin/sandbox
1
2
3
```
