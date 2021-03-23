package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	ctx := initializeContext()

	template, err := template.New("template.yml").Funcs(registerHelpers()).ParseFiles("template.yml")
	fmt.Print(err)
	str := bytes.NewBufferString("")
	f, err := os.Create("testOutput.yml")
	if err != nil {
		fmt.Print(err)
	}
	err = template.Execute(str, ctx)
	fmt.Print(err)

	f.WriteString(fmt.Sprintf("%v", str))
}
func initializeContext() map[string]interface{} {

	return map[string]interface{}{
		"variables": map[string]string{"registryUrl": "www.testRegistryUrl.com", "containerRegistryName": "TestContainerRegistery"},
		"inputs": map[string]interface{}{"httpApplicationRoutingDomain": true,
			"reuseACR": "false", "clusterName": "testCluster", "existingContainerRegistryId": 1024, "containerRegistryName": "Test-Container-Registery-Name"},
		"assets": map[string]string{"deploymentFile": "deploymentFile", "serviceFile": "serviceFile"},
		"env":    map[string]string{"REGION_ID": "divanshu_region"},
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
