package helpers

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	const userAccess = 0700

	outputPath := "../test-env/recursivehierarchy_output"
	exitCode := 0
	ok := os.MkdirAll(outputPath, userAccess)
	if ok == nil {
		exitCode = m.Run()
	} else {
		os.Exit(-1)
	}
	ok = os.RemoveAll(outputPath)
	if ok != nil {
		os.Exit(-1)
	}
	os.Exit(exitCode)
}

func TestHierarchyRecreation(test *testing.T) {

	testPath := "../test-env/recursivehierarchy"
	outputPath := "../test-env/recursivehierarchy_output"

	hierarchy := &RecursiveHierarchy{}
	hierarchy.CreateHierarchy(testPath, nil)

	if hierarchy.err != nil {
		test.Errorf("Hierarchy getting due to test requirements failed:\n %s",
			hierarchy.err.Error())
	}

	recreator := &HierarchyRecreator{}
	recreator.RecreateFolders(hierarchy, outputPath)

	if recreator.err != nil {
		test.Errorf("Hierarchy not recreated properly - error: %s\n", recreator.err.Error())
	}

	pathsToTest := []string{"folder1", "folderB", "folderB/folderA"}

	for counter := 0; counter < len(pathsToTest); counter++ {
		actualPath := outputPath + "/" + pathsToTest[counter]
		if _, err := os.Stat(actualPath); os.IsNotExist(err) {
			test.Errorf("Recreated path `%s` does not exist.", actualPath)
		}
	}
}

func TestHierarchyCopy(test *testing.T) {

	testPath := "../test-env/recursivehierarchy"
	outputPath := "../test-env/recursivehierarchy_output"

	hierarchy := &RecursiveHierarchy{}
	hierarchy.CreateHierarchy(testPath, nil)

	if hierarchy.err != nil {
		test.Errorf("Hierarchy getting due to test requirements failed:\n %s",
			hierarchy.err.Error())
	}

	recreator := &HierarchyRecreator{}
	recreator.RecreateFolders(hierarchy, outputPath)
	recreator.CopyFiles(hierarchy)

	elementsToTest := []string{"fileA.txt", "folder1/fileC.txt",
		"folderB/fileX.txt", "folderB/folderA/fileY.txt"}

	for counter := 0; counter < len(elementsToTest); counter++ {
		actualPath := outputPath + "/" + elementsToTest[counter]
		if _, err := os.Stat(actualPath); os.IsNotExist(err) {
			test.Errorf("Recreated file`%s` does not exist.", actualPath)
		}
	}
}
