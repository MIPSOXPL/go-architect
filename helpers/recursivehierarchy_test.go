package helpers

import (
	"testing"
)

func TestRecursiveCreateHierarchyNoErrors(test *testing.T) {
	hierarchy := &RecursiveHierarchy{}
	hierarchy.CreateHierarchy("../test-env/recursivehierarchy", nil)

	if hierarchy.err != nil {
		test.Errorf("Hierarchy test failed:\n %s", hierarchy.err.Error())
	} else if hierarchy.files == nil {
		test.Errorf("Hierarchy test failed. No files getted.")
	} else if hierarchy.folders == nil {
		test.Errorf("Hierarchy test failed. No folders getted.")
	}

}

func TestRecursiveCreateHierarchyErrors(test *testing.T) {
	hierarchy := &RecursiveHierarchy{}
	hierarchy.CreateHierarchy("../bad-path23131231231231", nil)

	if hierarchy.err == nil {
		test.Errorf("Hierarchy test failed - error expected and there is no error.")
	}

}

func TableIndexError(test *testing.T) {
	if recovered := recover(); recovered != nil {
		if recovered == "runtime error: index out of range" {
			test.Errorf("Panic - index out of range.")
		} else {
			test.Errorf("Unknown panic during test.")
		}
	}
}

func TestRecursiveHierarchyItself(test *testing.T) {
	defer TableIndexError(test)

	const foldersSizeExpected = 3
	const filesSizeExpected = 4

	hierarchy := &RecursiveHierarchy{}
	hierarchy.CreateHierarchy("../test-env/recursivehierarchy", nil)

	if hierarchy.err != nil {
		test.Errorf("Hierarchy test failed:\n %s", hierarchy.err.Error())
	}

	namesToTest := []string{"folder1", "folderB", "folderA", "fileA.txt", "fileC.txt",
		"fileX.txt", "fileY.txt"}

	for counter := 0; counter < foldersSizeExpected; counter++ {
		if hierarchy.folders[counter].GetName() != namesToTest[counter] {
			test.Errorf("There is no folder '%s' registered; folder '%s' expected.",
				namesToTest[counter], hierarchy.folders[counter].GetName())
		}
	}

	for counter := 0; counter < filesSizeExpected; counter++ {
		if hierarchy.files[counter].GetName() != namesToTest[foldersSizeExpected+counter] {
			test.Errorf("There is no file '%s' registered; file '%s' expected.",
				namesToTest[foldersSizeExpected+counter], hierarchy.files[counter].GetName())
		}
	}
}

func TestRecursiveCleanHierarchy(test *testing.T) {
	hierarchy := &RecursiveHierarchy{}
	hierarchy.CreateHierarchy("../test-env/recursivehierarchy", nil)

	if hierarchy.err != nil {
		test.Errorf("Hierarchy test failed:\n %s", hierarchy.err.Error())
	}

	hierarchy.CleanHierarchy()
	if hierarchy.err != nil || hierarchy.folders != nil || hierarchy.files != nil || hierarchy.sourcePath != "" {
		test.Errorf("Hierarchy not properly cleaned.")
	}
}
