package main

import (
	"fmt"
	"strings"
)

type CIPath struct {
	Platform string
	RepoType string
}

func cleanPlatform(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

func NewCIPath(platform string) CIPath {
	platformLower := cleanPlatform(platform)

	var repoType string
	switch platformLower {
	case "gitHub":
		repoType = ".github/workflows"
		platform = "Github"
	case "gitLab":
		repoType = ".gitlab-ci.yml"
		platform = "Gitlab"
	case "bitbucket":
		repoType = "bitbucket-pipelines.yml"
		platform = "Bitbucket"
	default:
		repoType = ""
		platform = "Unknown"
	}

	return CIPath{
		Platform: platform,
		RepoType: repoType,
	}
}

type Config struct {
	Appname string   `yaml:"app_name"`
	Port    int      `yaml:"port"`
	Steps   []string `yaml:"steps"`
}

func main() {

	ciPath := NewCIPath("GitHub")
	fmt.Printf("%v\n%v", ciPath.Platform, ciPath.RepoType)
	// 	data := Config{
	// 		Appname: "Inventory-service",
	// 		Port:    8080,
	// 		Steps:   []string{"Install", "Test", "Deploy"},
	// 	}

	// 	fmt.Println("hello world")

	// 	file, _ := os.Create("config.yml")
	// 	defer file.Close()

	// 	encoder := yaml.NewEncoder(file)
	// 	encoder.Encode(data)

}
