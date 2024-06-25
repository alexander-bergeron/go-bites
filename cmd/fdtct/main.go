package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

type TemplateType int

const (
	JSONTemplate TemplateType = iota
	YAMLTemplate
	TOMLTemplate
	TextTemplate
)

func determineTemplateType(input string) TemplateType {
	// Try to unmarshal as JSON
	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(input), &jsonData); err == nil {
		return JSONTemplate
	}

	// Try to unmarshal as YAML
	var yamlData map[string]interface{}
	if err := yaml.Unmarshal([]byte(input), &yamlData); err == nil {
		return YAMLTemplate
	}

	// Try to unmarshal as TOML
	var tomlData map[string]interface{}
	if _, err := toml.Decode(input, &tomlData); err == nil {
		return TOMLTemplate
	}

	// Assume it's text
	return TextTemplate
}

func isFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func checkFile(fileName string) {
	input := fileName
	if isFile(fileName) {
		fileContent, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		input = string(fileContent)
	}

	fmt.Println("Checking Content:")
	fmt.Printf("%q\n", input)

	// Test determineTemplateType function
	templateType := determineTemplateType(input)

	switch templateType {
	case JSONTemplate:
		fmt.Println("Type: JSON Template")
	case YAMLTemplate:
		fmt.Println("Type: YAML Template")
	case TOMLTemplate:
		fmt.Println("Type: TOML Template")
	case TextTemplate:
		fmt.Println("Type: Text Template")
	}
}

func main() {
	checkFile("example.yaml")
	checkFile("example.json")
	checkFile("example.toml")
	checkFile("example.txt")

	jsonInput := `{"field1": "value1", "field2": 123}`
	singleLineYamlInput := `field1: value1 field2: 123`
	multiLineYamlInput := `
field1: value1
field2: 123
`
	tomlInput := `
field1 = "value1"
field2 = 123
`

	checkFile(jsonInput)
	checkFile(singleLineYamlInput)
	checkFile(multiLineYamlInput)
	checkFile(tomlInput)
}
