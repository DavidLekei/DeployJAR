package main

import (
	"fmt"
	"os"

	ap "mpi.mb.ca/DeployJAR/argparser"
)

var config *map[string]string

func setConfigValue(setting string, value string) {
	(*config)[setting] = value
}

func parseArgs(args []string) {
	var options = make(map[string]ap.Option)

	options["-e"] = ap.Option{
		ConfigValue: "ENVIRONMENT",
		Required:    false,
	}

	options["-p"] = ap.Option{
		ConfigValue: "CURRENT_PI",
		Required:    false,
	}

	options["-s"] = ap.Option{
		ConfigValue: "SPRINT",
		Required:    false,
	}

	options["-j"] = ap.Option{
		ConfigValue: "JAR_NAME",
		Required:    false,
	}

	options["-v"] = ap.Option{
		ConfigValue: "VERBOSE",
		Required:    false,
	}

	parser := ap.New(options)

	parser.Parse(args, config)
}

func main() {

	config = LoadConfig()

	args := os.Args[1:]

	parseArgs(args)

	PrintConfig(config)

	environments := LoadEnvironments()

	environment := GetEnvironment((*config)["ENVIRONMENT"], environments)

	fmt.Println("Environment data: ", environment)

	fmt.Println("Done")
}
