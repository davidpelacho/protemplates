package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ansrivas/protemplates/golang"
	"github.com/ansrivas/protemplates/internal"
	"github.com/ansrivas/protemplates/python"
)

func handleGolangProject(projectName string) {
	var scm, username string

	fmt.Println("Please enter a desired scm eg. github.com or bitbucket.com")
	fmt.Scanf("%s", &scm)

	fmt.Println("Please enter a username corresponding to your scm eg. github.com/ansrivas, then ansrivas")
	fmt.Scanf("%s", &username)

	if username == "" || scm == "" {
		fmt.Println("Username or scm can not be empty")
		return
	}

	implementation := golang.Golang{Scm: scm, Username: username}
	err := internal.Create(implementation, projectName)
	if err != nil {
		panic(fmt.Sprintf("Unable to create the project: %s", projectName))
	}
	fmt.Printf("Successfully created golang project %s in current directory\n", projectName)

}

func handlePythonProject(projectName string) {
	implementation := python.Python{}
	err := internal.Create(implementation, projectName)
	if err != nil {
		panic(fmt.Sprintf("Unable to create the project: %s", projectName))
	}
	fmt.Printf("Successfully created python project %s in current directory\n", projectName)
}

func handleProjectCreation(language string) {
	lang := internal.SanitizeInput(language)

	var projectName string
	fmt.Println("Please enter a desired project name:")
	fmt.Scanf("%s", &projectName)
	switch strings.ToLower(lang) {
	case "python":
		handlePythonProject(projectName)
	case "go", "golang":
		handleGolangProject(projectName)
	default:
		fmt.Printf("\033[31mException: Language %s is currently not supported.\033[39m\n\n", language)
		os.Exit(1)
	}

}