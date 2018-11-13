package helpers

import (
	"io/ioutil"

	"github.com/MicroProcessingSolutions/go-architect/resource"
)

//ResursiveHierarchy contains folders and files resources detected during hierarchy detection on disk
type RecursiveHierarchy struct {
	folders    []resource.Folder
	files      []resource.File
	err        error
	sourcePath string
}

//CreateHierarchy creates hierarchy from given folder
func (hierarchy *RecursiveHierarchy) CreateHierarchy(path string, parent *resource.Folder) {
	hierarchy.sourcePath = path

	info, err := ioutil.ReadDir(path)

	if err != nil {
		hierarchy.err = err
	}

	for counter := range info {
		fullPath := path + "/" + info[counter].Name()
		if info[counter].IsDir() {
			folder := resource.Folder{Name: info[counter].Name(),
				Path:   fullPath,
				Parent: parent}
			hierarchy.folders = append(hierarchy.folders, folder)
			hierarchy.CreateHierarchy(fullPath, &hierarchy.folders[len(hierarchy.folders)-1])
		} else {
			file := resource.File{Name: info[counter].Name(),
				Path:   fullPath,
				Parent: parent}
			hierarchy.files = append(hierarchy.files, file)
		}
	}

	hierarchy.err = err
}
