### Go Module

#### some commands


```go build```
```go run```
```go mod help```

↳ geneally it writes git go model name like ``` go mod init github.com/bikashsaud/<reponame>```


### go mod uses
- initialise project
- give idea of 3rd party\
- also define which version of go lang

#### Versioning:

- Initial release:  1.0.0
- Patch release: 1.0.1
- Minor release: Major 1.1.0
- Major release: 2.0.0

### Commands

- ```go env``` : show all environments packages
- ```go mod tidy```: go module directions
- ```go mod verify```: to verify used pkg dependencies from go.sum
- ```go list```: list of all packages
- ```go list all```: list of all packages
- ```go list -m all```: list of all packages which are used or installed.
- ```go list -m version```: list of all
- 
- ```go mod why <pkg name>```: go mod why github.com/gorilla/mux --> to get which dependencies are usd from gorilla mux

- ```go mod graph```: to get the graph of dependencies
- ``` go mod edit -go 1.17```: edit go version 
- ``` go mod edit -module 1.10```: edit module version



### vendor: 
> ``` go mod vendor```
> it is another useful feature which get the all installs pkg to vendor folder.

- use to push to server and help to get dependencies from vendor not internet'
- first check vendor then check internet for this use cmd below
- ``` go run mod = vendor main.go```

