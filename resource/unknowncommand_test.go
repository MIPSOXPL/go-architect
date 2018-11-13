package resource

import "testing"

func TestUnknownCommandError(test *testing.T) {
	err := new(UnknownCommand)
	err.SetName("deploy")
	if err.Error() != "Unknown command: deploy.\n" {
		test.Errorf("UnknownCommand's Error failed.")
	}
}
