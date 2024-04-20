package main

import (
	"fmt"
	"os"
	"strings"

	io "mpi.mb.ca/DeployJAR/io"
)

var defaultConfigPath = "\\res\\config.cfg"

func PrintConfig(config *map[string]string) {
	fmt.Println("Config:")
	for key, value := range *config {
		fmt.Println("[" + key + "] : " + value)
	}
}

func LoadConfig() *map[string]string {
	return LoadConfigFromPath(defaultConfigPath)
}

func LoadConfigFromPath(configFilePath string) *map[string]string {
	dir, err := os.Getwd()
	file, err := os.Open(dir + configFilePath)
	defer file.Close()
	if err != nil {
		panic("ERROR: Could not open file: " + dir + configFilePath)
	}

	reader := io.LineReader{
		Offset: 0,
		File:   file,
	}

	config := make(map[string]string)

	line, err := reader.ReadLine()
	for err == nil {
		s := strings.Split(line, "=")
		config[s[0]] = strings.TrimSuffix(s[1], "\r")
		line, err = reader.ReadLine()
	}

	return &config
}
