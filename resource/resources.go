package resource

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Resources struct {
	name    string
	folders []Folder
}

func (resources *Resources) Save(filename string) error {
	return nil
}

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
