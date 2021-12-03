package cmd

var CommandCollection = make(map[string]Command)

type Command interface {
	BindConfig(interface{})
	Execute()
}
