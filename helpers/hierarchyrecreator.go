package helpers

import (
	"os"

	"github.com/MicroProcessingSolutions/go-architect/resource"
)

//HierarchyRecreator allows to recreate hierarchy
type HierarchyRecreator struct {
}

//Recreate allows to recreate hierarchy of hierarchy
func (recreator *HierarchyRecreator) Recreate(hierarchy *RecursiveHierarchy, path string) error {
	if hierarchy.err != nil {
		return hierarchy.err
	}

	return recreator.recreate(hierarchy, nil, path)
}

func (recreator *HierarchyRecreator) recreate(hierarchy *RecursiveHierarchy, parent *resource.Folder, path string) error {

	for counter := range hierarchy.folders {
		actualFolder := hierarchy.folders[counter].GetParent()

		if actualFolder.GetParent() == parent {
			preparedPath := path + "/" + actualFolder.GetName()
			err := os.Mkdir(preparedPath, os.ModeDir)

			if err != nil {
				return err
			}

			recreator.recreate(hierarchy, actualFolder, preparedPath)
		}
	}

	return nil
}
