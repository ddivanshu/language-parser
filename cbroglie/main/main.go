package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/cbroglie/mustache"
)

func main() {

	ctx := initializeContext()

	content, err := ioutil.ReadFile("template.yml")
	if err != nil {
		fmt.Print(err)
	}
	fileContent := string(content)
	template, err := mustache.Render(fileContent, ctx)
	f, err := os.Create("testOutput.yml")
	if err != nil {
		fmt.Print(err)
	}

	f.WriteString(fmt.Sprintf("%v", template))
}
func initializeContext() map[string]interface{} {

	return map[string]interface{}{
		"variables": map[string]string{"registryUrl": "www.testRegistryUrl.com", "containerRegistryName": "TestContainerRegistery"},
		"inputs": map[string]interface{}{"httpApplicationRoutingDomain": false,
			"reuseACR": "true", "clusterName": "testCluster", "existingContainerRegistryId": 1024, "containerRegistryName": "Test-Container-Registery-Name"},
		"assets":   map[string]string{"deploymentFile": "deploymentFile", "serviceFile": "serviceFile"},
		"env":      map[string]string{"REGION_ID": "divanshu_region"},
		"projects": []string{"p1", "p2", "p3", "p4"},
	}
}
func registerHelpers() template.FuncMap {
	return template.FuncMap{"equals": func(input string, expected string, csSensitive int) bool {
		return true
	}, "toLower": func(str string) string {
		return strings.ToLower(str)
	},
	}
}
