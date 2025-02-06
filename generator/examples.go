package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	examplesLogPrefix       = "[Generator Examples]"
	exampleTemplatesDirPath = "./templates/examples"
	// examplePath         = "../target/provider/provider_conf.json"
)

func GenerateExamples(envData map[string]string) {

	files := []string{}
	// Parse all templates recursively
	err := filepath.WalkDir(exampleTemplatesDirPath, func(path string, dirEntry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !dirEntry.IsDir() && filepath.Ext(path) == ".tmpl" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, filePath := range files {

		// Parse the template
		tmpl, err := template.ParseFiles(filePath)
		if err != nil {
			log.Fatal(err)
		}

		// Create the path to the output file
		newPath := strings.Replace(filePath, "templates", "../target", 1)
		newPath = strings.TrimSuffix(newPath, ".tmpl")

		// Create the parent directories
		err = os.MkdirAll(filepath.Dir(newPath), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		// Create the file
		file, err := os.Create(newPath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Execute the template
		err = tmpl.Execute(file, envData)
		if err != nil {
			log.Fatal(err)
		}
	}
}
