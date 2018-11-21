package resource

//Folder defines structure for folder
type Folder struct {
	Name   string   `json:"name"`
	Path   string   `json:"path"`
	Parent Resource `json:"parent"`
}

//GetName gets name of folder
func (folder *Folder) GetName() string {
	return folder.Name
}

//SetName sets name of folder
func (folder *Folder) SetName(name string) {
	folder.Name = name
}

//GetPath gets path of folder
func (folder *Folder) GetPath() string {
	return folder.Path
}

//SetPath sets path of folder
func (folder *Folder) SetPath(path string) {
	folder.Path = path
}

//GetParent gets parent of folder
func (folder *Folder) GetParent() Resource {
	return folder.Parent
}
