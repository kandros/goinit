# goinit

small utility to create new go projects inside GOPATH

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
