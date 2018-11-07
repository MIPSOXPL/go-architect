package tasks

import (
	"log"

	"github.com/MicroProcessingSolutions/go-architect/resource"
	"github.com/urfave/cli"
)

//AddHandler add resource to project
func AddHandler(c *cli.Context) error {
	resources, err := resource.LoadHandler()

	if err != nil {
		log.Fatal("Error occured during adding resource.")
		return nil
	}

	switch c.Command.Name {
	case "file":
		err = AddFileHandler(c, resources)
	case "folder":
		//TODO: implement
	default:
		commandError := new(resource.UnknownCommand)
		commandError.SetName(c.Command.Name)
		err = commandError
	}

	return resource.SaveHandler(resources)
}

//Add file handler adds file to resource
func AddFileHandler(c *cli.Context, resources *resource.Resources) error {
	nameArg := c.Args().Get(0)

	if nameArg != "" {
		fileResource := resource.File{Name: nameArg,
			Path:   nameArg,
			Parent: nil}
		resources.Files = append(resources.Files, fileResource)
	}

	return nil
}
