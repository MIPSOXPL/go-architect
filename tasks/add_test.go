package tasks

import (
	"testing"

	"github.com/MicroProcessingSolutions/go-architect/resource"
)

func TestAddFileHandler(test *testing.T) {
	resources := new(resource.Resources)
	resources.CreateResources("TestResources")

	if resources.Files != nil || len(resources.Files) > 0 {
		test.Errorf("Files array not empty.")
	}

	if AddFileHandler("123.txt", resources) != nil {
		test.Errorf("There was an error during adding file resource.")
	}

	if len(resources.Files) != 1 {
		test.Errorf("Bad files array length.")
	}

	if resources.Files[0].GetName() != "123.txt" {
		test.Errorf("File element not added properly.")
	}

	if AddFileHandler("player.png", resources) != nil {
		test.Errorf("There was an error during adding file resource.")
	}

	if len(resources.Files) != 2 {
		test.Errorf("Bad files array length.")
	}

	if resources.Files[1].GetName() != "player.png" {
		test.Errorf("File element not added properly.")
	}
}

func TestAddFolderHandler(test *testing.T) {
	resources := new(resource.Resources)
	resources.CreateResources("TestResources")

	if resources.Folders != nil || len(resources.Folders) > 0 {
		test.Errorf("Folders array not empty.")
	}

	if AddFolderHandler("data", resources) != nil {
		test.Errorf("There was an error during adding folder resource.")
	}

	if len(resources.Folders) != 1 {
		test.Errorf("Bad folders array length.")
	}

	if resources.Folders[0].GetName() != "data" {
		test.Errorf("Folder element not added properly.")
	}

	if AddFolderHandler("sprites", resources) != nil {
		test.Errorf("There was an error during adding folder resource.")
	}

	if len(resources.Folders) != 2 {
		test.Errorf("Bad folders array length.")
	}

	if resources.Folders[1].GetName() != "sprites" {
		test.Errorf("Folder element not added properly.")
	}
}
