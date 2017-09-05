package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func FromFile(filename string) (interface{}, error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return &Config{}, fmt.Errorf("can not open file: %v", err)
	}

	return fromReader(file)
}

func fromReader(r []byte) (*Config, error) {
	config := new(Config)
	err := json.Unmarshal(r, &config)

	if err != nil {
		return config, fmt.Errorf("can not parse config: %v", err)
	}

	return config, nil
}
