
** Micro Service **

- Get go-kit package

```bash
go get github.com/go-kit/kit
```

- Set GO env
- Run as follow
 
```bash
go run main.go
```

** Test **
```bash
curl -v -XPOST -d '{"name":"myname"}' http://172.17.0.2:8080/HelloWorld
```
