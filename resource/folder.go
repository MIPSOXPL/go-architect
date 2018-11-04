package resource

//Folder defines structure for folder
type Folder struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (folder *Folder) GetName() string {
	return folder.Name
}
func (folder *Folder) SetName(name string) {
	folder.Name = name
}
func (folder *Folder) GetPath() string {
	return folder.Path
}
func (folder *Folder) SetPath(path string) {
	folder.Path = path
}
