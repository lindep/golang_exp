Compile time variables
=====

Based on article
https://husobee.github.io/golang/compile/time/variables/2015/12/03/compile-time-const.html

Run go program with nm to print symbol table, i.e.

```bash
$ go build main.go
$ ./main
version:  

$ go tool nm prog | grep vers

4fa3b0 D main.version
447c40 T runtime.vdso_find_version
451990 T type..eq.runtime.version_key
451910 T type..hash.runtime.version_key
```

```bash
$ cd /homedir/dev_env/golang/src/github.com/lindep/golang_exp/compileTimeVars
$ go build -ldflags "-X main.version=$(git rev-parse master)" \
main.go
```
