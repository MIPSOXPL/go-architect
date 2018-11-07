package resource

//File contains information about file as resource in project.
type File struct {
	Name   string    `json:"name"`
	Path   string    `json:"path"`
	Parent *Resource `json:"parent"`
}
