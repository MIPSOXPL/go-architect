package helpers

import (
	"strings"
)

//PathResolver is used for resolving entire path of file/folder
type PathResolver struct {
	input  string
	output []string
	err    error
}

//Resolve is used for resolve path using input string
func (resolver *PathResolver) Resolve(input string) error {
	resolver.input = input
	countSlash := strings.Count(input, "/")
	countBackslash := strings.Count(input, "\\")

	if (countSlash > 0 && countBackslash > 0) || countSlash+countBackslash == len(resolver.input) {
		resolver.output = nil
		inconsistentError := BadResolving{}
		inconsistentError.SetPath(input)
		resolver.err = inconsistentError
		return inconsistentError
	} else if countSlash > 0 {
		resolver.output = strings.Split(resolver.input, "/")
	} else if countBackslash > 0 {
		resolver.output = strings.Split(resolver.input, "\\")
	} else {
		resolver.output = []string{input}
	}

	resolver.err = nil
	return nil
}

//CreatePathFromHierarchy is used for creating path
func (resolver *PathResolver) CreatePathFromHierarchy(hierarchy HierarchyRecreator) (string, error) {
	return "", nil
}
