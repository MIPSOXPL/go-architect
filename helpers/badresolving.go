package helpers

import (
	"fmt"
)

//BadResolving is error when something bad happens during resolving
type BadResolving struct {
	path string
}

//SetPath sets path which invoked error
func (err *BadResolving) SetPath(path string) {
	err.path = path
}

//Error implements error interface
func (err BadResolving) Error() string {
	if len(err.path) > 0 {
		return fmt.Sprintf("Path `%s` is bad.\n", err.path)
	}
	return fmt.Sprintf("Path is empty.\n")

}
