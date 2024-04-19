package argparser

type Option struct {
	Op       string
	Required bool
	Callback func(string)
}

type argparser struct {
	options map[string]Option
	//options []Option
}

func New(options map[string]Option) *argparser {
	ap := argparser{options: options}
	return &ap
}

func (parser argparser) Parse(args []string) {

	for index, arg := range args {
		_, ok := parser.options[arg]

		if ok {
			if index+1 < len(args) {
				parser.options[arg].Callback(args[index+1])
			} else {
				panic("Argument not supplied a value")
			}
		}
	}
}
