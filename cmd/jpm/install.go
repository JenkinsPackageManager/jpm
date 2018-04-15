package main

import (
	"log"
	"os"
)

// Install all dependencies defined in jpm.yml from jpm registry
func Install() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	createModulesDir()

	for _, dependency := range config.Dependencies {
		log.Println("Installing dependency " + dependency + "...")
		err = installDependency(dependency)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
