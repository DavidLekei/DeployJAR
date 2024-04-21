package main

import (
	"fmt"
	"os"

	ap "mpi.mb.ca/DeployJAR/argparser"
	io "mpi.mb.ca/DeployJAR/io"
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

func getEnvironmentFilePath(environment *Environment, serverIndex int, area string) string {
	path := "\\\\" + environment.Servers[serverIndex] + "\\"
	if area == "app" {
		return path + environment.AppLibDir
	}

	if area == "services" {
		return path + environment.ServicesLibDir
	}

	if area == "custom" {
		return path + environment.CustomServicesLibDir
	}

	return ""
}

func deployJar(config *map[string]string, environment *Environment) {
	jarPath := string((*config)["TFS_ROOT"]) + "\\PI" + (*config)["CURRENT_PI"] + "\\S" + (*config)["SPRINT"] + "\\" + (*config)["JAR_FOLDER"] + "\\" + (*config)["JAR_NAME"]

	if (*config)["VERBOSE"] == "true" {
		fmt.Println("JAR FILE: ", jarPath)
	}

	for index, _ := range environment.Servers {
		io.CopyFile(jarPath, getEnvironmentFilePath(environment, index, "app")+(*config)["JAR_NAME"])

		if (*config)["DEPLOY_TO_SERVICES"] == "true" {
			io.CopyFile(jarPath, getEnvironmentFilePath(environment, index, "services")+(*config)["JAR_NAME"])
			io.CopyFile(jarPath, getEnvironmentFilePath(environment, index, "custom")+(*config)["JAR_NAME"])
		}
	}

}

func main() {

	config = LoadConfig()

	args := os.Args[1:]

	parseArgs(args)

	environments := LoadEnvironments()

	environment := GetEnvironment((*config)["ENVIRONMENT"], environments)

	deployJar(config, environment)

	fmt.Println("Done")
}
