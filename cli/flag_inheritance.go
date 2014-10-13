package cli

// InheritFlags allow flag values inherit from the commands parent
func (cmd *CLI) InheritFlags(names ...string) {
	for _, name := range names {
		cmd.InheritFlag(name)
	}
}

// InheritFlag allows a flags value to inherit from the commands parent
func (cmd *CLI) InheritFlag(name string) {
	if cmd.parent == nil {
		panic("command does not have a parent")
	}
	flag := cmd.parent.(*CLI).getFlag(name)
	if cmd.inheritedFlags == nil {
		cmd.inheritedFlags = make(flagMap)
	}
	cmd.inheritedFlags[name] = flag
}

func (cmd *CLI) setFlagValuesFromParent() {
	for name, flag := range cmd.inheritedFlags {
		if _, exist := cmd.flags[name]; !exist {
			cmd.flagValues[flag] = cmd.parent.Flag(name)
		}
	}
}

// SubCommandsInheritFlags tells all subcommands to inherit flags
func (cmd *CLI) SubCommandsInheritFlags(names ...string) {
	for _, name := range names {
		cmd.SubCommandsInheritFlag(name)
	}
}

// SubCommandsInheritFlag tells all subcommands to inherit a flag
func (cmd *CLI) SubCommandsInheritFlag(name string) {
	flag := cmd.getFlag(name)
	if cmd.propogatingFlags == nil {
		cmd.propogatingFlags = make(flagMap)
	}
	cmd.propogatingFlags[name] = flag
}

func (cmd *CLI) copyPropogatingFlags() {
	if cmd.parent == nil {
		return
	}
	parentPropogatingFlags := cmd.parent.(*CLI).propogatingFlags
	cmd.propogatingFlags = parentPropogatingFlags.Without(cmd.flags).Merge(cmd.propogatingFlags)
	cmd.InheritFlags(parentPropogatingFlags.Names()...)
}
