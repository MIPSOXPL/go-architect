package resource

//File contains information about file as resource in project.
type File struct {
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	Parent *Folder `json:"parent"`
}

//GetName gets name of file
func (file *File) GetName() string {
	return file.Name
}

//SetName sets name of file
func (file *File) SetName(name string) {
	file.Name = name
}

//GetPath gets path of file
func (file *File) GetPath() string {
	return file.Path
}

//SetPath sets path of file
func (file *File) SetPath(path string) {
	file.Path = path
}

//GetParent gets parent of file
func (file *File) GetParent() *Folder {
	return file.Parent
}
