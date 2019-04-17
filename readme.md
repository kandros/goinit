# goinit

small utility to create new go projects inside GOPATH and open it using `$EDITOR`

## usage

`goinit <project_name>`

results in `<$GOPATH>/src/github.com/<org_name>/<project_name>/main.go` being created
with content

```go
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
```

> org_name is asked at first usage and value is stored in `~/.goinit/config.yaml` for next usages

## install

`go get -u github.com/kandros/goinit`

### flags

- `--no-open` will prevent editor opening
