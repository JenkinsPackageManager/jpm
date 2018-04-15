package main

import "github.com/JenkinsPackageManager/jpm-cli/util"

var (
	ConfigFileName = "jpm.yml"
	PackageZip     = "pkg.zip"
	ModulesDir     = "jpm_modules"
	CurrentDir     = util.GetCwd()
	Registry       = util.GetEnv("JPM_REGISTRY", "https://registry.jpm.org")
)
