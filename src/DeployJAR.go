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

func setPI(pi string) {

}

func setSprint(sprint string) {

}

func parseArgs(args []string) {
	var options = make(map[string]ap.Option)

	options["-e"] = ap.Option{
		Op:       "-e",
		Required: false,
		Callback: setEnvironment,
	}

	options["-p"] = ap.Option{
		Op:       "-p",
		Required: false,
		Callback: setPI,
	}

	options["-s"] = ap.Option{
		Op:       "-s",
		Required: false,
		Callback: setSprint,
	}

	parser := ap.New(options)

	parser.Parse(args)
}

func main() {

	config := LoadConfig()
	fmt.Println("config created: ", config)

	args := os.Args[1:]

	parseArgs(args)

	fmt.Println("Path to JAR File: ", filePath)
	fmt.Println("Environment: ", environment)

	fmt.Println("Done")
}
