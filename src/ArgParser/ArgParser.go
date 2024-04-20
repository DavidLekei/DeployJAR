package argparser

import (
	"fmt"
)

type Option struct {
	ConfigValue string
	Required    bool
}

type argparser struct {
	options map[string]Option
	//options []Option
}

func New(options map[string]Option) *argparser {
	ap := argparser{options: options}
	return &ap
}

func (parser argparser) Parse(args []string, config *map[string]string) {

	for index, arg := range args {
		_, ok := parser.options[arg]

		if ok {
			if index+1 < len(args) {
				fmt.Println("setting ", parser.options[arg].ConfigValue, " to : ", args[index+1])
				(*config)[parser.options[arg].ConfigValue] = args[index+1]
				// parser.options[arg].Callback(args[index+1])
			} else {
				panic("Argument not supplied a value")
			}
		}
	}
}
