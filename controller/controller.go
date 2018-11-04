//Package controller is package for controlling execution of CLI program
package controller

import (
	"github.com/urfave/cli"
)

//Controller is a controller for application managing
type Controller interface {
	SetApplication(application *cli.App, name string, usage string)
	PrepareApplication()
	Run()
}
