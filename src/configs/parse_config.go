package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func FromFile(filename string, object interface{}) error {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("can not open file: %v", err)
	}

	return fromReader(file, &object)
}

func fromReader(r []byte, object *interface{}) error {
	err := json.Unmarshal(r, &object)

	if err != nil {
		return fmt.Errorf("can not parse config: %v", err)
	}

	return nil
}
