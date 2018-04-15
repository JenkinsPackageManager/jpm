package main

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/JenkinsPackageManager/jpm-cli/util"
)

func getConfig() (Config, error) {
	config := Config{}

	err := util.ReadYAML(path.Join(CurrentDir, ConfigFileName), &config)

	return config, err
}

func writeConfig(c Config) error {
	return util.WriteYAML(path.Join(CurrentDir, ConfigFileName), c)
}

func createModulesDir() error {
	return os.Mkdir(path.Join(CurrentDir, ModulesDir), 0777)
}

func installDependency(dependency string) error {
	name, version := parseDependency(dependency)

	return fetchDependency(name, version)
}

func parseDependency(dependency string) (string, string) {
	parts := strings.Split(dependency, "@")

	name := parts[0]
	version := "latest"

	if len(parts) == 2 {
		version = parts[1]
	}

	return name, version
}

func fetchDependency(name, version string) error {
	dir := path.Join(CurrentDir, ModulesDir, name)
	pkg := path.Join(dir, PackageZip)
	err := os.Mkdir(dir, 0777)
	if err != nil {
		return err
	}

	log.Println("http " + Registry + "/package/" + name + "@" + version)
	err = util.FetchFile(Registry+"/package/"+name+"@"+version, pkg)
	if err != nil {
		return err
	}

	err = extractDependency(pkg)
	if err != nil {
		return err
	}

	return err
}

func extractDependency(dir string) error {
	return nil
}
