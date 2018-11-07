package resource

//LoadHandler loads project file
func LoadHandler() (*Resources, error) {
	loadedResources := new(Resources)
	err := loadedResources.Load("resources.json")
	return loadedResources, err
}
