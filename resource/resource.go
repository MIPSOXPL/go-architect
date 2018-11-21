package resource

//Resource defines resource to pack
type Resource interface {
	GetName() string
	SetName(string)
	GetPath() string
	SetPath(string)
	GetParent() Resource
}
