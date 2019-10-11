package loader

import (
	"encoding/json"
	"io/ioutil"
	"os"

	model "github.com/xsami/xgonet/models"
)

// UFStruct is the structure to parse the mock file with friend and user model
type UFStruct struct {
	Users   []model.User   `json:"users"`
	Friends []model.Friend `json:"friends"`
}

// LoadModel does what it say, open a json file and parse a file with the given structure
func LoadModel(filePath string, model *UFStruct) error {

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &model)
	if err != nil {
		return err
	}

	return nil
}
