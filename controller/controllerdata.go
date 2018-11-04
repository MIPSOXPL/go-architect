package controller

import (
	"log"
	"os"

	"github.com/MicroProcessingSolutions/go-architect/resource"
	"github.com/urfave/cli"
)

//Data contains data to manage CLI application
type Data struct {
	applicationHandle *cli.App
}

//SetApplication sets application for data
func (data *Data) SetApplication(application *cli.App, name string, usage string) {
	data.applicationHandle = application
	if data.applicationHandle != nil {
		data.applicationHandle.Name = name
		data.applicationHandle.Usage = usage
	}
}

//AddCommand adds command to application
func (data *Data) AddCommand(command cli.Command) {
	if data.applicationHandle != nil {
		commands := &data.applicationHandle.Commands
		*commands = append(*commands, command)
	}
}

//Run runs application
func (data *Data) Run() {
	if data.applicationHandle != nil {
		err := data.applicationHandle.Run(os.Args)

		if err != nil {
			log.Fatal(err)
		}
	}
}

//PrepareApplication prepares application for work
func (data *Data) PrepareApplication() {
	createCommand := cli.Command{Name: "create",
		Aliases: []string{"c"},
		Usage:   "Create resource file for new project",
		Action:  resource.SaveHandler,
	}
	data.AddCommand(createCommand)
}
