package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

const yamlSep = "\n---\n"

// panicf is fmt.Sprintf + panic
func panicf(s string, args ...interface{}) {
	panic(fmt.Sprintf(s, args...))
}

// sortYAML deserializes s to struct and then back to string returning sorted version.
func sortYAML(s string) string {
	m := make(map[string]interface{})
	if err := yaml.Unmarshal([]byte(s), &m); err != nil {
		panicf("Failed to unmarshal YAML: %v", err)
	}

	d, err := yaml.Marshal(m)
	if err != nil {
		panicf("Failed to marshal YAML: %v", err)
	}

	return string(d)
}

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panicf("Failed to read STDIN: %v", err)
	}

	// Sort every document.
	docs := strings.Split(string(in), yamlSep)
	sortedYAMLs := make([]string, len(docs))
	for i, doc := range docs {
		sortedYAMLs[i] = sortYAML(doc)
	}

	fmt.Print(strings.Join(sortedYAMLs, yamlSep))
}
