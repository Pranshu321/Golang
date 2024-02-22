### For all the version of any package

```bash
go list -m -versions github.com/gorilla/mux
github.com/gorilla/mux v1.2.0 v1.3.0 v1.4.0 v1.5.0 v1.6.0 v1.6.1 v1.6.2 v1.7.0 v1.7.1 v1.7.2 v1.7.3 v1.7.4 v1.8.0 v1.8.1
```
### For resolve dependency work in tidy mainly indirect flag remove

```bash
go mod tidy
```

### It will tell why i am using this package

```bash
go mod why github.com/gorilla/mux
```

### It will tell the graph format of why

```bash
go mod graph
```

### This will basically like node_modules of the golang and tell which pakcgae we are using

```bash
go mod vendor
```