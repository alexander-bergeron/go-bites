package main

import (
	"fmt"
	"regexp"
)

func formatString(template string, values map[string]string) string {
	re := regexp.MustCompile(`\{(\w+)\}`)
	return re.ReplaceAllStringFunc(template, func(match string) string {
		key := match[1 : len(match)-1]
		if val, ok := values[key]; ok {
			return val
		}
		return match
	})
}

func main() {
	template := "this is a {foo} and {bar}"
	values := map[string]string{
		"foo": "also test",
		"bar": "test",
	}

	result := formatString(template, values)
	fmt.Println(result) // Output: this is a also test and test
}
