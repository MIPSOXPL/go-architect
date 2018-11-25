package helpers

import "github.com/MicroProcessingSolutions/go-architect/resource"

//HierarchyRecreator allows to recreate hierarchy
type HierarchyRecreator struct {
}

//Recreate allows to recreate hierarchy of hierarchy
func (recreator *HierarchyRecreator) Recreate(hierarchy *RecursiveHierarchy, path string) error {
	if hierarchy.err != nil {
		return hierarchy.err
	}

	for {
		for {

		}
	}

	return nil
}

func (recreator *HierarchyRecreator) recreate(hierarchy *RecursiveHierarchy, parent *resource.Folder, path string) error {

	/*for counter := range hierarchy.folders {
		//actualFolder := hierarchy.folders[counter].GetParent()
	}*/

	return nil
}
