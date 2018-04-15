package main

import "log"

func Add(dependency string) {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	newDeps := make([]string, len(config.Dependencies)+1)

	createModulesDir()

	installDependency(dependency)

	newDeps = append(config.Dependencies, dependency)
	config.Dependencies = newDeps

	writeConfig(config)
}
