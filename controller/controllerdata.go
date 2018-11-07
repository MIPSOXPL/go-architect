package controller

import (
	"log"
	"os"

	"github.com/MicroProcessingSolutions/go-architect/resource"
	"github.com/MicroProcessingSolutions/go-architect/tasks"
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
		Action:  resource.CreateHandler,
	}
	buildCommand := cli.Command{Name: "build",
		Aliases: []string{"b"},
		Usage:   "Build project using go-build",
		Action:  tasks.BuildHandler,
	}

	folderCommand := cli.Command{Name: "folder",
		Aliases:   []string{"fo"},
		Usage:     "Folder as input for resource",
		ArgsUsage: "First argument is folder name/path",
		Action:    tasks.AddHandler,
	}

	fileCommand := cli.Command{Name: "file",
		Aliases:   []string{"fi"},
		Usage:     "File as input for resource",
		ArgsUsage: "First argument is file name/path",
		Action:    tasks.AddHandler,
	}

	addCommand := cli.Command{Name: "add",
		Aliases:     []string{"a"},
		Usage:       "Add resource to project",
		Subcommands: cli.Commands{folderCommand, fileCommand},
	}

	data.AddCommand(createCommand)
	data.AddCommand(buildCommand)
	data.AddCommand(addCommand)
}
