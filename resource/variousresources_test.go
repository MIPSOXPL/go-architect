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

}
func TestGetPath(test *testing.T) {

}
func TestSetPath(test *testing.T) {

}
func TestGetParent(test *testing.T) {

}
