package resource

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

//CreateHandler saves resource from CLI context
func CreateHandler(c *cli.Context) error {
	resource := new(Resources)
	actualFile, err := os.Getwd()

	if err != nil {
		return nil
	}

	resource.CreateResources(filepath.Base(actualFile))

	return resource.Save("resources.json")
}
