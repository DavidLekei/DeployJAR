package main

import (
	"fmt"
	"os"

	ap "mpi.mb.ca/DeployJAR/argparser"
)

var filePath string
var environment string

func setFilePath(path string) {
	filePath = path
}

func setEnvironment(env string) {
	environment = env
}

func parseArgs(args []string) {
	var options = make(map[string]ap.Option)

	options["-f"] = ap.Option{
		Op:       "-f",
		Required: true,
		Callback: setFilePath,
	}

	options["-e"] = ap.Option{
		Op:       "-e",
		Required: true,
		Callback: setEnvironment,
	}

	parser := ap.New(options)

	parser.Parse(args)
}

func main() {
	args := os.Args[1:]

	parseArgs(args)

	fmt.Println("Path to JAR File: ", filePath)
	fmt.Println("Environment: ", environment)

	fmt.Println("Done")
}
