package tasks

import "testing"

func TestUnknownCommandError(test *testing.T) {
	err := new(BadFilename)
	err.SetFilename("")
	if err.Error() != "Filename `` is bad.\n" {
		test.Errorf("UnknownCommand's Error failed.")
	}

	err.SetFilename("&&_A")
	if err.Error() != "Filename `&&_A` is bad.\n" {
		test.Errorf("UnknownCommand's Error failed.")
	}
}
