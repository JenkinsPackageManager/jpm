package util

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ReadYAML reads a file and maps its contents to a given model struct
func ReadYAML(path string, model interface{}) error {
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlFile, model)
}

// WriteYAML outputs the contents of data structure into a yaml file
func WriteYAML(path string, data interface{}) error {
	contents, _ := yaml.Marshal(&data)

	return ioutil.WriteFile(path, contents, 0644)
}
