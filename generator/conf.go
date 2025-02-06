package main

import (
	"log"
	"os"
	"text/template"
)

var (
	confLogPrefix    = "[Generator Provider Conf]"
	ConfTemplatePath = "./templates/provider_conf_template.json.tmpl"
	ConfPath         = "../target/provider/provider_conf.json"
)

func GenerateProviderConf(envData map[string]string) {

	tmpl, err := template.ParseFiles(ConfTemplatePath)
	if err != nil {
		log.Println(confLogPrefix+" failed to parse files: ", err)
		return
	}

	outputFile, err := os.Create(ConfPath)
	if err != nil {
		log.Println(confLogPrefix+" failed to create output file: ", err)
		return
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, envData)
	if err != nil {
		log.Println(confLogPrefix+" failed to execute template: ", err)
		return
	}
}
