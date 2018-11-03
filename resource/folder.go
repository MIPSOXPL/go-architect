package resource

//Folder defines structure for folder
type Folder struct {
	name string
	path string
}

func (folder *Folder) GetName() string {
	return folder.name
}
func (folder *Folder) SetName(name string) {
	folder.name = name
}
func (folder *Folder) GetPath() string {
	return folder.path
}
func (folder *Folder) SetPath(path string) {
	folder.path = path
}
