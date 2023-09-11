package executor

import (
	"log"
)

type Command struct {
	Name        string
	Run         func() error
	subCommands []*Command
}

func (c *Command) AddCommand(cmd *Command) error {
	if c == nil || cmd == nil {
		return nil
	}
	if c.Name == cmd.Name {
		log.Panicf("child command cannot have the same name as their parent: %s", cmd.Name)
		return nil
	}
	c.subCommands = append(c.subCommands, cmd)

	return nil
}

func Execute(cmd *Command) error {
	return execute(cmd, true)
}

func execute(cmd *Command, root bool) error {
	// log.Printf("%+v", cmd)

	if len(cmd.subCommands) == 0 {
		// log.Print("# " + cmd.Name + " running")
		go run(cmd)
	}

	for _, subCommand := range cmd.subCommands {
		// log.Print("# " + subCommand.Name + " running")
		go execute(subCommand, false)
	}

	if root {
		// log.Print("# " + cmd.Name + " running")
		run(cmd)
	}

	return nil
}

func run(cmd *Command) error {
	log.Print("# " + cmd.Name + " running")
	return cmd.Run()
}
