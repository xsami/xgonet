package main

import "encoding/json"

func LoadModel(data []byte, model interface{}) error {

	err := json.Unmarshal([]byte(data), &model)

	return err

}
