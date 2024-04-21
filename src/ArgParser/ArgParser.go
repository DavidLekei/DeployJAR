package argparser

import (
	"fmt"
	"os"
)

type Option struct {
	ConfigValue string
	Required    bool
	Help        string
}

type argparser struct {
	options map[string]Option
	//options []Option
}

func New(options map[string]Option) *argparser {
	ap := argparser{options: options}
	return &ap
}

func (parser argparser) Help() {
	fmt.Println("DeployJAR - Deploys the MPI.jar to a specified environments WAR file.")
	fmt.Println("\n\tNOTE: You can set values in res/config.cfg to avoid using command line arguments every time\n")
	fmt.Println("The following options can be used to override values found in res/config.cfg")
	fmt.Println("Options:")

	for _, value := range parser.options {
		fmt.Println(value.Help)
	}

	os.Exit(0)
}

func (parser argparser) Parse(args []string, config *map[string]string) {

	for index, arg := range args {
		//Check for "--help" first, since it causes an exit
		if arg == "--help" {
			parser.Help()
		}

		_, ok := parser.options[arg]

		if ok {
			if index+1 < len(args) {
				//If the NEXT argument starts with a - , then it's assumed to be another argument, and hence the current argument was not supplied a value
				if args[index+1][0] == '-' {
					panic("Argument [" + args[index] + "] not supplied a value")
				}
				(*config)[parser.options[arg].ConfigValue] = args[index+1]
			} else {
				panic("Argument [" + args[index] + "] not supplied a value")
			}
		}
	}
}
