package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Error: missing project name \n\nusage: goinit <project_name>\n\n")
		os.Exit(0)
	}

	gopath := os.Getenv("GOPATH")
	projectName := "gino"
	username := "kandros"
	orgPath := "src/github.com/" + username
	fileName := "main.go"
	projectPath := path.Join(gopath, orgPath, projectName)
	filePath := path.Join(projectPath, fileName)

	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		os.MkdirAll(projectPath, 0777)
		err := ioutil.WriteFile(filePath, []byte(mainFileContent), 0777)

		if err != nil {
			panic(err)
		}
		fmt.Printf("created project at %s", filePath)

	} else {
		fmt.Printf("file at path %s already exists\n", filePath)
	}

}

const mainFileContent = `
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
`
