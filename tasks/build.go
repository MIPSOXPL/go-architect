//Package tasks is used for creating tasks that are related to go programs
package tasks

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli"
)

//Build is a struct for build command
type Build struct {
	output     string
	buildError error
}

//BuildProject builds actual project
func (build *Build) BuildProject() string {
	output, err := exec.Command("go", "build").CombinedOutput()
	build.buildError = err
	build.output = string(output)

	if err != nil {
		return fmt.Sprintf("go-build not executed properly, trace:\n%s", build.output)
	}

	if len(build.output) == 0 {
		return "go-build executed succesfully"
	}
	return fmt.Sprintf("go-build trace:\n%s", build.output)
}

//BuildHandler builds project from CLI
func BuildHandler(c *cli.Context) error {

	log.Println("Project build invoked!")
	build := new(Build)
	log.Println(build.BuildProject())

	return build.buildError
}
