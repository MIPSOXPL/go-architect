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
		err = AddFileHandler(c.Args().Get(0), resources)
	case "folder":
		err = AddFolderHandler(c.Args().Get(0), resources)
	default:
		commandError := new(resource.UnknownCommand)
		commandError.SetName(c.Command.Name)
		err = commandError
	}

	return resource.SaveHandler(resources)
}

//AddFileHandler handler adds file to resources
func AddFileHandler(nameArg string, resources *resource.Resources) error {

	if nameArg != "" {
		fileResource := resource.File{Name: nameArg,
			Path:   nameArg,
			Parent: nil}
		resources.Files = append(resources.Files, fileResource)
	}

	return nil
}

//AddFolderHandler adds folder to resources
func AddFolderHandler(nameArg string, resources *resource.Resources) error {
	if nameArg != "" {
		folderResource := resource.Folder{Name: nameArg,
			Path:   nameArg,
			Parent: nil}
		resources.Folders = append(resources.Folders, folderResource)
	} else {

	}

	return nil
}
