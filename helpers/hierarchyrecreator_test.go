package helpers

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	outputPath := "../test-env/recursivehierarchy_output"
	exitCode := 0
	ok := os.MkdirAll(outputPath, os.ModeDir)
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
	recreator.Recreate(hierarchy, outputPath)

	pathsToTest := []string{"folder1", "folderB", "folderB/folderA"}

	for counter := 0; counter < len(pathsToTest); counter++ {
		actualPath := outputPath + "/" + pathsToTest[counter]
		if _, err := os.Stat(actualPath); os.IsNotExist(err) {
			test.Errorf("Recreated path `%s` does not exist.", actualPath)
		}
	}
}
