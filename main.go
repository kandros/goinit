package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	gomain "github.com/kandros/gomain/pkg"
	"github.com/kandros/goutil/editorutil"
	"github.com/spf13/viper"
)

func main() {
	home := os.Getenv("HOME")
	viper.SetConfigFile(home + "/.goinit/config.yaml")
	viper.SetDefault("orgname", "")
	viper.SetDefault("open_in_editor", true)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, err := os.Stat(home + "/.goinit/config.yaml"); os.IsNotExist(err) {
			os.MkdirAll(home+"/.goinit", 0777)
			os.Create(home + "/.goinit/config.yaml")
		}
	}

	for _, a := range os.Args {
		if a == "--no-open" {
			viper.Set("open_in_editor", false)
		}
	}

	var args []string

	for _, val := range os.Args[1:] {
		if strings.HasPrefix(val, "--") || strings.HasPrefix(val, "-") {
			continue
		} else {
			args = append(args, val)
		}
	}

	if len(args) == 0 {
		var projectname string
		fmt.Print("project name: ")
		fmt.Scan(&projectname)
		viper.Set("projectname", projectname)
	} else {
		viper.Set("projectname", args[0])
	}

	if viper.GetString("orgname") == "" {
		var orgname string
		fmt.Print("orgname: ")
		fmt.Scan(&orgname)
		viper.Set("orgname", orgname)
		err := viper.WriteConfig()
		if err != nil {
			panic(err)
		}
	}

	gopath := os.Getenv("GOPATH")
	projectname := viper.GetString("projectname")
	orgname := viper.GetString("orgname")
	orgPath := "src/github.com/" + orgname
	fileName := "main.go"
	projectPath := path.Join(gopath, orgPath, projectname)
	filePath := path.Join(projectPath, fileName)

	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		fmt.Printf("file at path %s already exists\n", filePath)
		os.Exit(0)
	}

	os.MkdirAll(projectPath, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Printf("created project at %s\n", filePath)

	err = os.Chdir(projectPath)
	if err != nil {
		panic(err)
	}

	gomain.CreateMainFile()
	gomain.CreateMainTestFile()
	if viper.GetBool("open_in_editor") {
		editorutil.OpenProjectInEditor(projectPath, filePath)
	}

}
