package helpers

import "testing"

func TestPathResolverResolveSlash(test *testing.T) {
	resolver := new(PathResolver)
	err := resolver.Resolve("abc/def/geh")

	expectedValues := []string{"abc", "def", "geh"}

	if resolver.input != "abc/def/geh" {
		test.Errorf("Resolver input not saved.")
	}

	if err != nil {
		test.Errorf("Error was returned, even if there is no error during resolve.")
	}

	if err != resolver.err {
		test.Errorf("Internal resolver error is not the same as returned one.")
	}

	if len(resolver.output) != len(expectedValues) {
		test.Errorf("Bad size of resolved output. Expected: %d; Provided: %d.",
			len(expectedValues), len(resolver.output))
	} else {
		for counter := range resolver.output {
			if resolver.output[counter] != expectedValues[counter] {
				test.Errorf("Resolved strings vary. Expected: %s; Provided: %s.",
					expectedValues[counter], resolver.output[counter])
			}
		}
	}
}

func TestPathResolverResolveBackslash(test *testing.T) {
	resolver := new(PathResolver)
	err := resolver.Resolve("abc\\def\\geh")

	expectedValues := []string{"abc", "def", "geh"}

	if resolver.input != "abc\\def\\geh" {
		test.Errorf("Resolver input not saved.")
	}

	if err != nil {
		test.Errorf("Error was returned, even if there is no error during resolve.")
	}

	if err != resolver.err {
		test.Errorf("Internal resolver error is not the same as returned one.")
	}

	if len(resolver.output) != len(expectedValues) {
		test.Errorf("Bad size of resolved output. Expected: %d; Provided: %d.",
			len(expectedValues), len(resolver.output))
	} else {
		for counter := range resolver.output {
			if resolver.output[counter] != expectedValues[counter] {
				test.Errorf("Resolved strings vary. Expected: %s; Provided: %s.",
					expectedValues[counter], resolver.output[counter])
			}
		}
	}
}

func TestPathResolverResolveOnlyName(test *testing.T) {
	resolver := new(PathResolver)
	err := resolver.Resolve("data")

	expectedValues := []string{"data"}

	if resolver.input != "data" {
		test.Errorf("Resolver input not saved.")
	}

	if err != nil {
		test.Errorf("Error was returned, even if there is no error during resolve.")
	}

	if err != resolver.err {
		test.Errorf("Internal resolver error is not the same as returned one.")
	}

	if len(resolver.output) != len(expectedValues) {
		test.Errorf("Bad size of resolved output. Expected: %d; Provided: %d.",
			len(expectedValues), len(resolver.output))
	} else {
		for counter := range resolver.output {
			if resolver.output[counter] != expectedValues[counter] {
				test.Errorf("Resolved strings vary. Expected: %s; Provided: %s.",
					expectedValues[counter], resolver.output[counter])
			}
		}
	}
}

func TestPathResolverResolveEmpty(test *testing.T) {
	resolver := new(PathResolver)
	err := resolver.Resolve("")

	if resolver.input != "" {
		test.Errorf("Resolver input not saved.")
	}

	if len(resolver.output) > 0 {
		test.Errorf("Output should be empty.")
	}

	if err != resolver.err {
		test.Errorf("Internal resolver error is not the same as returned one.")
	}

}

func TestPathResolverResolveEmptySlashes(test *testing.T) {
	resolver := new(PathResolver)
	err := resolver.Resolve("///")

	if resolver.input != "///" {
		test.Errorf("Resolver input not saved.")
	}

	if len(resolver.output) > 0 {
		test.Errorf("Output should be empty.")
	}

	if err != resolver.err {
		test.Errorf("Internal resolver error is not the same as returned one.")
	}

}

func TestPathResolverResolveBadPath(test *testing.T) {
	resolver := new(PathResolver)
	err := resolver.Resolve("abc/def\\geh")

	if resolver.input != "abc/def\\geh" {
		test.Errorf("Resolver input not saved.")
	}

	if len(resolver.output) != 0 {
		test.Errorf("Something was resolved even with error.")
	}

	if err == nil {
		test.Errorf("Error was not returned, even if there is error during resolve.")
	}

	if err.Error() != "Path `abc/def\\geh` is bad.\n" {
		test.Errorf("Bad error message.")
	}

	if err != resolver.err {
		test.Errorf("Internal resolver error is not the same as returned one.")
	}
}

func TestPathResolverResolveEmptyPath(test *testing.T) {
	resolver := new(PathResolver)
	err := resolver.Resolve("")

	if resolver.input != "" {
		test.Errorf("Resolver input not saved.")
	}

	if len(resolver.output) != 0 {
		test.Errorf("Something was resolved even with error.")
	}

	if err == nil {
		test.Errorf("Error was not returned, even if there is error during resolve.")
	}

	if err.Error() != "Path is empty.\n" {
		test.Errorf("Bad error message.")
	}

	if err != resolver.err {
		test.Errorf("Internal resolver error is not the same as returned one.")
	}
}
