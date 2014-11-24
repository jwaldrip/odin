package cli

import "fmt"

type subCommandList []*subCommandListItem
type subCommandListItem struct {
	name    string
	command *SubCommand
}

// Each loops through each subcommand
func (l *subCommandList) Each(fn func(string, *SubCommand)) {
	for _, item := range *l {
		fn(item.name, item.command)
	}
}

func (l *subCommandList) Add(name string, command *SubCommand) {
	listItem := &subCommandListItem{name: name, command: command}
	*l = append(*l, listItem)
}

func (l *subCommandList) Get(name string) (*SubCommand, error) {
	var subCommand *SubCommand
	for _, item := range *l {
		if item.name == name {
			subCommand = item.command
			break
		}
		if item.command.nameAliases[name] != "" {
			subCommand = item.command
			break
		}
	}
	if subCommand == nil {
		return nil, fmt.Errorf("invalid subcommand")
	}
	return subCommand, nil
}
