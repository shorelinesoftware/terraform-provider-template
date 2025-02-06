// Copyright 2021, Shoreline Software Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"os"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate tofu fmt -recursive ../target/examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs --provider-dir=../target --website-source-dir=../generator/templates/doc_templates --provider-name=terraform-provider-$PROVIDER_SHORT_NAME --rendered-provider-name=$RENDERED_PROVIDER_NAME --tf-version=1.5.7

var (
	mainLogPrefix string = "[ MAIN ]"
)

func main() {

	providerBrand := os.Getenv("PROVIDER_BRAND")
	useLocal := os.Getenv("USE_LOCAL")

	envData, err := GetEnvData(providerBrand, useLocal)
	if err != nil {
		log.Println(mainLogPrefix+" failed to get env data: ", err)
		return
	}

	GenerateProviderConf(envData)
	GenerateExamples(envData)

}
