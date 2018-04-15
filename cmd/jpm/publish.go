package main

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/JenkinsPackageManager/jpm-cli/util"
)

// Publish some package to jpm registry
func Publish() {
	// check jmp.yml exists
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// add files to zip
	outFile := path.Join(CurrentDir, PackageZip)

	err = toZip(CurrentDir, outFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Println("zip √")
	}

	// send file
	err = util.PostFile(Registry+"/package/"+config.Name+"@"+config.Version, outFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Println("upload √")
	}

	// remove zip file
	err = os.Remove(outFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Println("clean √")
	}

	log.Println("Published " + config.Name + "@" + config.Version)
}

func toZip(directory, file string) error {
	out, err := os.Create(file)
	defer out.Close()

	if err != nil {
		return err
	}

	pkg := zip.NewWriter(out)
	defer pkg.Close()

	return addFiles(pkg, directory, "")
}

func addFiles(pkg *zip.Writer, basePath, baseInZip string) error {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && file.Name() != PackageZip {
			content, err := ioutil.ReadFile(path.Join(basePath, file.Name()))
			if err != nil {
				return err
			}

			f, err := pkg.Create(path.Join(baseInZip + file.Name()))
			if err != nil {
				return err
			}
			_, err = f.Write(content)
			if err != nil {
				return err
			}
		} else if file.IsDir() {
			newBase := path.Join(basePath, file.Name())
			addFiles(pkg, newBase, file.Name())
		}
	}

	return err
}
