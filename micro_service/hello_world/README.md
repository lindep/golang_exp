
* Micro Service *

** get go-kit **

```bash
go get github.com/go-kit/kit
```

** Run as follow **
Set GO env
```bash
go run main.go
```

** Test **
```bash
curl -v -XPOST -d '{"name":"myname"}' http://172.17.0.2:8080/HelloWorld
```
