package resource

import "testing"

func TestGetName(test *testing.T) {
	var file Resource
	file = &File{Name: "abcdefghij"}
	var folder Resource
	folder = &Folder{Name: "123"}

	if file.GetName() != "abcdefghij" {
		test.Errorf("File resource GetName failed.")
	}
	if folder.GetName() != "123" {
		test.Errorf("File resource GetName failed.")
	}
}
func TestSetName(test *testing.T) {
	var file Resource
	file = &File{}
	file.SetName("abcdefghij")
	var folder Resource
	folder = &Folder{}
	folder.SetName("123")

	if file.GetName() != "abcdefghij" {
		test.Errorf("File resource SetName failed.")
	}
	if folder.GetName() != "123" {
		test.Errorf("File resource SetName failed.")
	}
}
func TestGetPath(test *testing.T) {
	var file Resource
	file = &File{Path: "/home/dev/abcdefghij"}
	var folder Resource
	folder = &Folder{Path: "C:\\Users\\Dev\\123"}

	if file.GetPath() != "/home/dev/abcdefghij" {
		test.Errorf("File resource GetPath failed.")
	}
	if folder.GetPath() != "C:\\Users\\Dev\\123" {
		test.Errorf("File resource GetPath failed.")
	}
}
func TestSetPath(test *testing.T) {
	var file Resource
	file = &File{}
	file.SetPath("/home/dev/abcdefghij")
	var folder Resource
	folder = &Folder{}
	folder.SetPath("C:\\Users\\Dev\\123")

	if file.GetPath() != "/home/dev/abcdefghij" {
		test.Errorf("File resource SetPath failed.")
	}
	if folder.GetPath() != "C:\\Users\\Dev\\123" {
		test.Errorf("File resource SetPath failed.")
	}
}
func TestGetParent(test *testing.T) {
	var folder Resource
	folder = &Folder{Parent: nil}
	var file Resource
	file = &File{Parent: folder}

	if folder.GetParent() != nil {
		test.Errorf("Folder resource parent not nil")
	}
	if file.GetParent() != folder {
		test.Errorf("File resource parent not folder")
	}
}
