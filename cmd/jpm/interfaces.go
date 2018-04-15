package main

// Config describes the format of jpm configuration file
type Config struct {
	Name         string   `yaml:"name"`
	Version      string   `yaml:"version"`
	Main         string   `yaml:"main"`
	Files        []string `yaml:"files"`
	Dependencies []string `yaml:"dependencies"`
}
