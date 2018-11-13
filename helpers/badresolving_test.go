package helpers

import "testing"

func TestBadResolvingError(test *testing.T) {
	err := new(BadResolving)
	err.SetPath("/abcd\\efgh/ijk")
	if err.Error() != "Path `/abcd\\efgh/ijk` is bad.\n" {
		test.Errorf("BadResolving's Error failed.")
	}

	err.SetPath("///")
	if err.Error() != "Path `///` is bad.\n" {
		test.Errorf("BadResolving's Error failed.")
	}

	err.SetPath("")
	if err.Error() != "Path is empty.\n" {
		test.Errorf("BadResolving's Error failed.")
	}
}
