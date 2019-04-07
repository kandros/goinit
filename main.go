package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Error: missing project name \n\nusage: goinit <project_name>\n\n")
		os.Exit(0)
	}

	gopath := os.Getenv("GOPATH")
	projectName := args[0]
	username := "kandros"
	orgPath := "src/github.com/" + username
	fileName := "main.go"
	projectPath := path.Join(gopath, orgPath, projectName)
	filePath := path.Join(projectPath, fileName)

	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		fmt.Printf("file at path %s already exists\n", filePath)
		os.Exit(0)
	}

	os.MkdirAll(projectPath, 0777)
	err := ioutil.WriteFile(filePath, []byte(mainFileContent), 0777)

	if err != nil {
		panic(err)
	}
	fmt.Printf("created project at %s", filePath)

	for _, a := range args {
		if a == "--open" {
			editor := os.Getenv("EDITOR")

			var cmd *exec.Cmd
			if editor == "code" || editor == "code-insider" {
				cmd = exec.Command(editor, projectPath, "--goto", filePath)
			} else {
				cmd = exec.Command(editor, projectPath)
			}

			cmd.Start()
		}
	}

}

const mainFileContent = `
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
`
