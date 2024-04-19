package main

import (
	"encoding/json"
	"fmt"
)

var defaultEnvFilePath = "C:/ADO/"

type Environment struct {
	name                 string
	servers              []string
	appLibDir            string
	servicesLibDir       string
	customServicesLibDir string
}

type FileReader struct {
}

func (reader FileReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func LoadEnvironments(envFile string) []*Environment {

	var reader FileReader

	decoder := json.NewDecoder(reader)
	fmt.Println("decoder created: ", decoder)

	return nil
}

func GetEnvironment(envName string) *Environment {
	return nil
}
