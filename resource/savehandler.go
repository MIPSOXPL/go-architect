package resource

//SaveHandler saves resource from CLI context
func SaveHandler(resources *Resources) error {
	return resources.Save("resources.json")
}
