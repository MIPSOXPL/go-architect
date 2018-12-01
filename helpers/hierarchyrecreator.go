package helpers

import (
	"os"

	"errors"

	"github.com/MicroProcessingSolutions/go-architect/resource"
)

//HierarchyRecreator allows to recreate hierarchy
type HierarchyRecreator struct {
	savedPath          string
	hierarchyGenerated bool
	err                error
	originHierarchy    *RecursiveHierarchy
}

//RecreateFolders allows to recreate hierarchy of hierarchy
func (recreator *HierarchyRecreator) RecreateFolders(hierarchy *RecursiveHierarchy, path string) error {
	recreator.hierarchyGenerated = false
	if hierarchy.err != nil {
		recreator.err = hierarchy.err
		return hierarchy.err
	}

	recreator.err = recreator.recreate(hierarchy, nil, path)

	if recreator.err == nil {
		recreator.savedPath = path
		recreator.hierarchyGenerated = true
	}

	return recreator.err
}

func (recreator *HierarchyRecreator) recreate(hierarchy *RecursiveHierarchy, parent *resource.Folder, path string) error {

	const userAccess = 0700

	for counter := range hierarchy.folders {
		actualFolder := hierarchy.folders[counter]
		actualParent := actualFolder.GetParent()
		parentEqual := false
		if actualParent == parent {
			parentEqual = true
		} else if actualParent != nil && parent != nil {
			parentEqual = *actualParent == *parent
		}
		if parentEqual {
			preparedPath := path + "/" + actualFolder.GetName()
			err := os.Mkdir(preparedPath, userAccess)

			if err != nil {
				recreator.err = err
				return err
			}

			recreator.recreate(hierarchy, &actualFolder, preparedPath)
		}
	}

	return nil
}

//CopyFiles copies files after recreation of structure of folders
func (recreator *HierarchyRecreator) CopyFiles(hierarchy *RecursiveHierarchy) error {
	if recreator.originHierarchy != hierarchy {
		recreator.err = errors.New("Origin hierarchy not same as argument")
		return recreator.err
	}
	if recreator.hierarchyGenerated == false {
		recreator.err = errors.New("Hierarchy folders not generated earlier")
		return recreator.err
	}

	copiedFolders := hierarchy.folders
	copiedFiles := hierarchy.files

	return nil
}
