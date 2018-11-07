package resource

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//Resources contains resources for project
type Resources struct {
	Name    string   `json:"name"`
	Folders []Folder `json:"folders"`
	Files   []File   `json:"file"`
}

//CreateResources creates resources internals
func (resources *Resources) CreateResources(name string) {
	resources.Name = name
}

//Save saves resources file to json
func (resources *Resources) Save(filename string) error {

	data, err := json.Marshal(resources)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}

//Load loads resources file from json
func (resources *Resources) Load(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, resources)

	if err != nil {
		return err
	}

	return nil
}
