package main

import "github.com/joho/godotenv"

var (
	NggImportFile       string = "./env/release/ngg.env"
	ShorelineImportFile string = "./env/release/shoreline.env"
	LocalDevImportFile  string = "./env/dev/local.env"
	VariablesImportFile string = "./env/dev/variables.env"
)

func GetVariablesEnv() (envData map[string]string, err error) {
	return godotenv.Read(VariablesImportFile)
}

func GetLocalDevEnv() (envData map[string]string, err error) {
	return godotenv.Read(LocalDevImportFile)
}

func GetReleaseEnv(providerBrand string) (envData map[string]string, err error) {
	return godotenv.Read(GetImportFileName(providerBrand))
}

func GetImportFileName(providerBrand string) (importFileName string) {
	if providerBrand == "ngg" {
		importFileName = NggImportFile
	} else {
		importFileName = ShorelineImportFile
	}
	return importFileName
}

func GetEnvData(providerBrand string, useLocal string) (envData map[string]string, err error) {

	if useLocal != "" {
		GenerateLocalDevEnv(providerBrand)
		envData, err = GetLocalDevEnv()
	} else {
		envData, err = GetReleaseEnv(providerBrand)
	}

	return envData, err
}
