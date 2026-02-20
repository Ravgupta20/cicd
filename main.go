package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Workflow struct {
	Name string         `yaml:"name"`
	On   []string       `yaml:"on"`
	Jobs map[string]Job `yaml:"jobs"`
}

type Job struct {
	RunsOn string `yaml:"runs-on"`
	Steps  []Step `yaml:"steps"`
}

type Step struct {
	Name string `yaml:"name,omitempty"`
	Uses string `yaml:"uses,omitempty"`
	Run  string `yaml:"run,omitempty"`
}

func main() {
	var filePath = "./.github/workflows/"
	var file *os.File

	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		fmt.Printf("error creating directory: %v\n", err)
	}

	workflow := Workflow{
		Name: "Basic CI",
		On:   []string{"push"},
		Jobs: map[string]Job{
			"build": {
				RunsOn: "ubuntu-latest",
				Steps: []Step{
					{
						Name: "Checkout code",
						Uses: "actions/checkout@v3",
					},
					{
						Name: "Install dependencies",
						Run:  "go mod download"},
					{
						Name: "Run tests",
						Run:  "go test ./...",
					},
				},
			},
		},
	}

	file, err = os.Create(filePath + "main.yml")
	if err != nil {
		fmt.Printf("error creating file: %v\n", err)
	}

	defer file.Close()
	encoder := yaml.NewEncoder(file)
	encoder.Encode(workflow)

}
