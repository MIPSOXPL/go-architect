package resource

import (
	"fmt"
)

//UnknownCommand represents error of unknown command situation
type UnknownCommand struct {
	name string
}

//SetName for command that is unknown
func (commandError *UnknownCommand) SetName(name string) {
	commandError.name = name
}

func (commandError *UnknownCommand) Error() string {
	return fmt.Sprintf("Unknown command: %s.", commandError.name)
}
