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

> org_name is asked at first usage and store in `~/.goinit/config.yaml`

### flags

- `—-org <org_name>` will use a specific org

  - if the org_name provided is not yet saved it will promt to create and save it

- `--no-open` will prevent editor opening
