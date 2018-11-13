package tasks

import (
	"fmt"
)

//BadFilename is error of filename that is bad
type BadFilename struct {
	name string
}

//SetFilename sets name of filename which invoked error
func (err *BadFilename) SetFilename(name string) {
	err.name = name
}

//Error implements error interface
func (err *BadFilename) Error() string {
	return fmt.Sprintf("Filename `%s` is bad.\n", err.name)
}
