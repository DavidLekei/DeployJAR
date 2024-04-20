package main

import (
	"encoding/json"
	"io"
	"os"
)

var defaultEnvFilePath = "C:/ADO/"

type Environment struct {
	Name                 string   `json:"name"`
	Servers              []string `json:"servers"`
	AppLibDir            string   `json:"appLibDir"`
	ServicesLibDir       string   `json:"servicesLibDir"`
	CustomServicesLibDir string   `json:"customServicesLibDir"`
}

type Environments struct {
	Environments []Environment `json:"environments"`
}

func LoadEnvironments() *Environments {
	return LoadEnvironmentsWithPath("\\res\\environments.json")
}

func LoadEnvironmentsWithPath(envFile string) *Environments {
	var environments Environments

	dir, err := os.Getwd()
	file, err := os.Open(dir + envFile)

	defer file.Close()

	if err != nil {
		panic("ERROR - environments.json file not found")
	}

	bytes, err := io.ReadAll(file)

	if err != nil {
		panic("ERROR - Unable to READ environments.json")
	}

	json.Unmarshal(bytes, &environments)

	return &environments
}

func GetEnvironment(env string, environments *Environments) *Environment {
	for _, environment := range environments.Environments {
		if environment.Name == env {
			return &environment
		}
	}
	return nil
}
